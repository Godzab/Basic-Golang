package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var err error

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

var db *sql.DB

func init() {
	db, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
		return
	}
	err = db.Ping()
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Connected to DB")
}

func main() {
	http.HandleFunc("/bookstore", index)
	http.HandleFunc("/book", bookShow)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "This route only gets.", http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM books")
	check(err, w)
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		check(err, w)
		bks = append(bks, bk)
	}

	check(rows.Err(), w)

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}

func bookShow(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "This route only gets.", http.StatusMethodNotAllowed)
		return
	}

	isbn := req.FormValue("isbn")

	if isbn == "" {
		http.Error(w, "Isbn required.", 400)
		return
	}
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := Book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
	case err != nil:
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
}

func check(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, "This route only gets.", http.StatusMethodNotAllowed)
		return
	}
}

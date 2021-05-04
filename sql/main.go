package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/gotest?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/files", http.FileServer(http.Dir(".")))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully Completed")
	check(err)
}

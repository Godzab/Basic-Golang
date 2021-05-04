package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/files", bar)
	http.HandleFunc("/rdt", rdt)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	dt := req.FormValue("q")
	io.WriteString(w, "Do my search: "+dt)
}

func rdt(w http.ResponseWriter, req *http.Request) {
	fmt.Println("You current method at redirect is: ", req.Method)
	http.Redirect(w, req, "/files", http.StatusMovedPermanently)
}

func bar(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		// open the file
		fl, h, err := req.FormFile("q")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("\n FILE: ", fl, "\n HEADER: ", h, "\n ERR: ", err)

		//read
		bs, err := ioutil.ReadAll(fl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="POST" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
		</form>
	`+s)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "visits",
		Value: "1",
	})
	io.WriteString(w, "The cookie was written")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Godhiza")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "The cookie was found: ", c)
}

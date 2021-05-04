package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var tpl *template.Template
var fn = template.FuncMap{}

func init() {
	tpl = template.Must(template.ParseFiles("parsedForm.gohtml"))
}

func main() {
	http.HandleFunc("/", b)
	http.HandleFunc("/animals/", c)
	http.ListenAndServe(":8080", nil)
}

func c(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("API-KEY", "123456789")
	w.Header().Set("Content-Type", "text/html; charset utf-8")
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	data := struct {
		Method      string
		submissions url.Values
	}{
		r.Method,
		r.Form,
	}
	tpl.ExecuteTemplate(w, "parsedForm.gohtml", data)
}

func b(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Snoop doggy dog")
	case "/cat":
		io.WriteString(w, "Pussy cat")
	}
}

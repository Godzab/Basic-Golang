package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dg", lfile)
	http.ListenAndServe(":8080", nil)
}

func lfile(w http.ResponseWriter, r *http.Request) {
	fl, err := os.Open("./main.go")

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer fl.Close()

	io.Copy(w, fl)
}

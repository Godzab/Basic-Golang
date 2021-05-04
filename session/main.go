package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var m = map[string]string{}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewString()
		c = &http.Cookie{
			Name:  "session",
			Value: id,
		}
		http.SetCookie(w, c)
	}
	pwd, _ := bcrypt.GenerateFromPassword([]byte("Nyandoro23"), bcrypt.MinCost)
	fmt.Println(string(pwd))
	fmt.Println(c)
}

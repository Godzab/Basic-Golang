package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type person struct {
	Name    string
	Surname string
	Actions []string
}

type cities struct {
	Name      string  `json: "name"`
	Longitude float64 `json: "lng"`
	Latitude  float64 `json: "lat"`
}

func main() {
	c := getHmacCode("test@gmail.com")
	fmt.Println(c)
	c = getHmacCode("test@gmail.com")
	fmt.Println(c)

	d := b64Encode("Godzila GODHINGTA *#%$&")
	fmt.Println(d)

	http.HandleFunc("/", ctx)
	http.HandleFunc("/enc", enc)
	http.HandleFunc("/msh", enc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// Creates one way encryption - Value->harsh
func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte("harsh_key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func b64Encode(s string) string {
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	return s64
}

func ctx(w http.ResponseWriter, req *http.Request) {
	cx := req.Context()
	log.Println(cx)
	fmt.Fprintln(w, cx)
}

func enc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Nice", "Ice", "things"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}

func mashl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Nice", "Ice", "things"},
	}

	jsn, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	data := person{}
	ss := string(jsn)
	err = json.Unmarshal([]byte(ss), data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
	w.Write(jsn)
}

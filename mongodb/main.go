package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"pitchperion.com/mongodb/controllers"
)

var ctl *controllers.UserController

func init() {
	ctl = controllers.NewUserController(getSession())
}

func main() {
	r := httprouter.New()
	r.GET("/users", ctl.Index)
	r.GET("/user/:id", ctl.GetUser)
	r.POST("/users", ctl.CreateUser)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost/myFirstDatabase")
	if err != nil {
		panic(err)
	}
	return s
}

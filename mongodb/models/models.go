package models

import "gopkg.in/mgo.v2/bson"

type Person struct {
	Name    string `json: "name" bson: "name"`
	Surname string `json: "surname" bson: "surname"`
	Age     int32  `json: "age" bson: "age"`
	Id      bson.ObjectId
}

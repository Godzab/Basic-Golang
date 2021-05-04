package main

import "fmt"

type contactInfo struct {
	email   string
	address string
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	var alex person
	alex.firstname = "Alex"
	alex.lastname = "Anderson"
	alex.updateName("Jim")
	alex.details()
}

func (pointerToPerson *person) updateName(nFirstName string) {
	(*pointerToPerson).firstname = nFirstName
}

func (p person) details() {
	fmt.Printf("%+v", p)
}

package main

import "fmt"

type bot interface {
	getGreeting() string
	fish(int) int
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	//eb := englishBot{}
	//sb := spanishBot{}

	//printGreeting(eb)
	//printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// Very custom logic here
	return "Hi there."
}

func (spanishBot) getGreeting() string {
	return "Hola amigo."
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

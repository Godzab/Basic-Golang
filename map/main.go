package main

import "fmt"

func main() {
	colors := map[string]string{
		"white": "#ffff",
		"red":   "#f10000",
		"green": "#f23123",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex for color", color, "is", hex)
	}
}

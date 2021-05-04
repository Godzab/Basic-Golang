package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("Failed to create server: ", err)
		os.Exit(1)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := "<!DOCTYPE html><html><head><title>Template</title></head><body><h1>IT WORKS</h1></body></html>"
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

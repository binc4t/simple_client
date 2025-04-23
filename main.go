package main

import (
	"log"
	"net"
)

func main() {
	// dial
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

}

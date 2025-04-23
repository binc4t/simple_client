package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// dial
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// send http request: /mock
	conn.Write([]byte("GET /mock HTTP/1.1\r\nHost: localhost:8080\r\n\r\n"))

	// read 1Bps
	buf := make([]byte, 1)

	tc, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("conn is not a *net.TCPConn")
	}
	err = tc.SetReadBuffer(1)
	if err != nil {
		log.Fatal(err)
	}

	size := 0
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		size += n
		fmt.Printf("read %d bytes\n", size)
		time.Sleep(1 * time.Second)
	}
}

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	io.Copy(os.Stdout, c)
}

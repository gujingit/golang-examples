package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	client, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err.Error())
	}
	defer client.Close()
	mustCopy(os.Stdout, client)
}

func mustCopy(dst *os.File, src net.Conn) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		conn.Close()
	}()

	done := make(chan int)
	go func() {
		mustCopy(os.Stdout, conn)
		log.Print("done")
		done <- 1
	}()
	mustCopy(conn, os.Stdin)
	//conn.(*net.TCPConn).CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

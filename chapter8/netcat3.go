package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		conn.Close()
	}()
	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer func() {
			log.Print("done")
			wg.Done()
		}()
		mustCopy(os.Stdout, conn)
	}()

	mustCopy(conn, os.Stdin)
	wg.Wait()
	conn.(*net.TCPConn).CloseWrite()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

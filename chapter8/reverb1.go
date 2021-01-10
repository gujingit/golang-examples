package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Printf("shout: %s \n", shout)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handle(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()

}

func main() {
	listener, e := net.Listen("tcp", "localhost:8080")
	if e != nil {
		log.Fatal(e)
	}

	defer func() {
		listener.Close()
	}()

	for {
		conn, _ := listener.Accept()
		handle(conn)
	}

}

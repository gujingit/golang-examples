package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	msg := make(chan struct{})
	go func() {
		for {
			if input.Scan() {
				go echo(c, input.Text(), 1*time.Second)
				msg <- struct{}{}
			}
		}

	}()

	ticker := time.NewTicker(1 * time.Second)
	timeout := 0
	for {
		select {
		case <-msg:
			timeout = 0
		default:
			<-ticker.C
			timeout++
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		if timeout > 10 {
			break
		}
	}

	ticker.Stop()
	c.Close()
	fmt.Println("close connection")
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Printf("shout: %s \n", shout)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
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
		handleConn(conn)
	}

}

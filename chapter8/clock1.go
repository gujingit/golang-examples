package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err.Error())
	}

	for {
		conn, _ := server.Accept()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "connection error: %s", err.Error())
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	//for {
	//	_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
	//	if err != nil {
	//		_, _ = fmt.Fprintf(os.Stdout, "err: %s", err.Error())
	//		break
	//	}
	//	time.Sleep(1 * time.Second)
	//}
	time.Sleep(2 * time.Minute)
}

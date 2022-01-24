package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var port = flag.Int("port", 8000, "clock server port")
var location = flag.String("location", "Asia/Shanghai", "location")

func main() {
	flag.Parse()
	server, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
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
	var cst, _ = time.LoadLocation(*location)
	for {
		_, err := io.WriteString(conn, time.Now().In(cst).Format("15:04:05\n"))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "err: %s", err.Error())
			break
		}
		time.Sleep(1 * time.Second)
	}
}

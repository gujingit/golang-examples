package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	clockWall := sync.Map{}

	for _, arg := range os.Args[1:] {
		attr := strings.Split(arg, "=")
		if len(attr) != 2 {
			fmt.Printf("split %s error\n", arg)
			continue
		}
		region := attr[0]
		port, err := strconv.Atoi(attr[1])
		if err != nil {
			fmt.Printf("conv str %s to int, error: %s \n", attr[1], err.Error())
			continue
		}
		go recvMsg(region, port, &clockWall)
	}

	for {
		clockWall.Range(func(key, value interface{}) bool {
			fmt.Printf("Region: %v\tTime:%v\n", key, value)
			return true
		})
		fmt.Println("--------------------")
		time.Sleep(3 * time.Second)
	}

}

func recvMsg(region string, port int, clockWall *sync.Map) {
	client, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	for {
		data := make([]byte, 1024*10)
		_, err = client.Read(data)
		if err != nil {
			fmt.Printf("read error: %s", err.Error())
			break
		}
		clockWall.Store(region, string(data))
	}
	_ = client.Close()
}

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case t := <-time.After(10 * time.Second):
		fmt.Println(t.Format("2006-01-02 15:04:05"))
	case <-abort:
		fmt.Println("cancel")
		return
	}

	launch()
}

func launch() {
	fmt.Println("fire!")
}

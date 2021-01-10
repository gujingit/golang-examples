package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	_ "os"
	"time"
)

func v1(args []string) {
	ch := make(chan string)
	start := time.Now()
	for _, url := range args {
		go fetch(url, ch)
	}

	for range args {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("failed to get url %s, err: %s", url, err.Error())
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("failed to read body from url %s, error: %s", url, err.Error())
		return
	}

	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), nbytes, url)
}

func main() {
	args := []string{"https://golang.org", "http://gopl.io", "https://godoc.org"}
	v1(args)
}

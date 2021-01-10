package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func v1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "error to get response from %s", url)
			os.Exit(1)
		}
		content, err := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "error to read response from %s", url)
			os.Exit(1)
		}
		fmt.Printf("response: %s \n", content)

	}
}

func v2() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "error to get response from %s", url)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "error to read response from %s", url)
			os.Exit(1)
		}
	}
}

func v3() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Printf("url: %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "error to get response from %s", url)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "error to read response from %s", url)
			os.Exit(1)
		}
	}
}

func v4() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "fail to get %s, error: %s \n", url, err.Error())
			os.Exit(1)
		}

		// print result
		 _, err = io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()
		 if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "fail to read response body, error: %s", err.Error())
			os.Exit(1)
		}

		// print http status
		_, _ = fmt.Fprintf(os.Stdout, "\n http status: proto %s, status %s", resp.Proto, resp.Status)
	}
}

func main() {
	v4()
}

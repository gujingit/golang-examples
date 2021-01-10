package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	duplicatedLines := make(map[string]int)

	for _, arg := range os.Args[1:] {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "open file err: %s", err.Error())
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			duplicatedLines[line]++
		}
	}

	for k, v := range duplicatedLines {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}

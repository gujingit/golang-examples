package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commandLines := make(map[string]int)
	// read from os stdin
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "" {
			break
		}
		commandLines[input.Text()]++
	}

	for key, value := range commandLines {
		if value > 1 {
			fmt.Printf("%d\t%s\n", value, key)
		}
	}
}

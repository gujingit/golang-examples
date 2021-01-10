// read from files
package main

import (
	"bufio"
	"fmt"
	"os"
)

//  go run dup2.go ./data.txt
func main() {
	duplicatedLines := make(map[string]int)
	for _, arg := range os.Args[1:] {

		file, err := os.Open(arg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "open file error: %s\n", err.Error())
			continue
		}
		commandLine(duplicatedLines, file)
		file.Close()
	}

	for k, v := range duplicatedLines {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}

func commandLine(lines map[string]int, file *os.File) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		lines[input.Text()]++
	}
}

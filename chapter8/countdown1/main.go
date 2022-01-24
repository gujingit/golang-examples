package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-ticker
	}

	launch()
}

func launch() {
	fmt.Println("fire!")
}

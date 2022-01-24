package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)
	go counter(in)
	go square(out,in)
	printer(out)

}

// 单向channel，只写
func counter(in chan<- int) {
	for i := 0; i < 100; i++ {
		in <- i
	}
	// 只有在发送者所在的goroutine才会调用close函数
	// 因此对一个只接收的channel调用close将是一个编译错误
	close(in)
}

func square(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)

}

func printer(out <-chan int) {
	for ret := range out{
		fmt.Printf("%d\n",ret)
	}

}

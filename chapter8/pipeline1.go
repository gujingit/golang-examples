package main

import "fmt"

func v1() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		// 如果设置i上限，则会进入死锁
		/*
			fatal error: all goroutines are asleep - deadlock!
		*/
		for i := 0; ; i++ {
			naturals <- i
		}
	}()

	go func() {
		// 从channel中读取元素时，需要使用for循环
		for {
			num := <-naturals
			squares <- num * num
		}

	}()

	for {
		fmt.Printf("%d\n", <-squares)
	}
}

// 生成有限序列
func v2() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				close(squares)
				break
			}
			squares <- x * x
		}
	}()

	for {
		ret, ok := <-squares
		if !ok {
			break
		}
		fmt.Printf("%d\n", ret)
	}

}

// 通过range遍历channel
func v3() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for ret := range squares {
		fmt.Printf("%d\n", ret)
	}

}

func main() {
	v3()
}

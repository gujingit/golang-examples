package main

import (
	"fmt"
	"sync"
)

func main() {
	// buffer 为1的阻塞chan
	sizes := make(chan int)
	var wg sync.WaitGroup // number of working goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// worker
		go func(i int) {
			// 在sizes的元素被读出来之前wg.Done()不会被执行
			defer wg.Done()
			// 阻塞，直到sizes里的元素被读出来
			sizes <- i
		}(i)
	}

	// 如果放到main线程里，无人消费sizes chan，会导致wg.Done()无法执行，wg.Wait()一直无法为0
	/*
		wg.Wait()
		close(sizes)
	*/

	/*
		正确写法
	*/
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int
	// range sizes，当channel被关闭并且没有值可接收时跳出循环
	for size := range sizes {
		total += size
	}
	fmt.Printf("total: %d", total)

	// 如果放到loop的后面，由于sizes channel一直没有关闭，因此无法结束循环，死锁。
	/*
		wg.Wait()
		close(sizes)
	*/

}

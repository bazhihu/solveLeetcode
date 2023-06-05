package main

// 十个生产者 十个消费者 通过

import (
	"fmt"
	"sync"
)

var sy = &sync.WaitGroup{}

func func1(t int, aChan chan int64) {
	var i int64 = 0
	for {
		if i > 10 {
			break
		}
		sy.Add(1)
		b := int64(t*100) + i
		aChan <- b
		i++
	}
}

func func2(aChan chan int64) {
	for t := range aChan {
		defer sy.Done()
		fmt.Println("resp", t)
	}
}

func main() {
	var (
		aChan = make(chan int64)
	)
	defer close(aChan)

	// 消费
	for j := 0; j < 10; j++ {
		go func2(aChan)
	}

	// 生产
	for i := 0; i < 10; i++ {
		go func1(i, aChan)
	}

	sy.Wait()
	aChan <- 1
}

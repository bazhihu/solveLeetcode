package main

/**
交替输出自然数 先输出 奇数 然后输出 偶数
*/

import (
	"fmt"
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer close(ch1)
	defer close(ch2)

	wg.Add(2)
	go say1(wg, ch1, ch2)
	go say2(wg, ch2, ch1)

	wg.Wait()

	str := "hello"
	var bs []byte = []byte(str)
	bs[0] = 'x'

	fmt.Println(string(bs))
}

func say1(wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch1 <- (i+1)*2 - 1
		fmt.Println(<-ch2)
	}
	return
}

func say2(wg *sync.WaitGroup, ch2 chan int, ch1 chan int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch1)
		ch2 <- (i + 1) * 2
	}
	return

}

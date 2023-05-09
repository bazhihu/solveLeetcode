package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("END")
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func demoWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

}

func main() {
	demoWithCancel()
	time.Sleep(3 * time.Second)
}

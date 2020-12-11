package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {

	f1, err := os.Create("./cpu.pprof")
	if err != nil {
		return
	}
	pprof.StartCPUProfile(f1)

	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 10; i++ {
		go logicCode()
	}

	wg.Wait()

	defer func() {
		pprof.StopCPUProfile()
		f1.Close()
	}()

	return
}

// 一段有问题的代码
func logicCode() {
	var c chan int // nil
	for {
		select {
		case v := <-c: // 阻塞
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			time.Sleep(time.Millisecond * 500)
			return
		}
	}
}

package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	//f1, err := os.Create("./cpu.pprof")
	//if err != nil {
	//	return
	//}
	//pprof.StartCPUProfile(f1)

	f2, err := os.Create("./mem.pprof")
	if err != nil {
		return
	}
	pprof.WriteHeapProfile(f2)

	for i := 0; i < 10; i++ {
		go logicCode()
	}

	time.Sleep(20 * time.Second)

	defer func() {
		pprof.StopCPUProfile()
		//f1.Close()
		f2.Close()
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
		}
	}
}

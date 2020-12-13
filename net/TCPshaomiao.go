package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	//_, err := net.Dial("tcp", "api.ikongji.com:22")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("Connection successful")
	//}
	//return

	f1, err := os.Create("./mem.pprof") //在当前路径下创建一个cpu.pprof 文件
	if err != nil {
		return
	}
	pprof.WriteHeapProfile(f1)

	hostName := flag.String("hostname", "www.baidu.com", "hostname to test")
	startPost := flag.Int("start-port", 1, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 22222, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Microsecond*200, "timeout")
	flag.Parse()

	ports := []int{}
	ch := make(chan int, 10)
	wg := &sync.WaitGroup{}
	wwg := &sync.WaitGroup{}

	// 遍历端口
	for port := *startPost; port < *endPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			if ok, _ := isOpen(*hostName, port, *timeout); ok {
				ch <- port
			}
		}(port)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// 从channel 中获取数据
	wwg.Add(1)
	go func() {
		defer wwg.Done()
		for {
			if c, ok := <-ch; ok {
				ports = append(ports, c)
			} else {
				break
			}
		}
		return
	}()

	wwg.Wait()

	fmt.Printf("opened ports: %v\n", ports)

	defer func() {
		pprof.StopCPUProfile()
		f1.Close()
	}()
}

func isOpen(host string, port int, timeout time.Duration) (bool, error) {

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true, err
	}

	return false, err
}

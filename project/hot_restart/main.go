package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

// 热重启
// 新老进程无缝切换，保持对Client的服务
// 原理
// 1、父进程监听重启信号
// 2、收到信号，父进程调用fork,同时传递socket 描述符给子进程
// 3、子进程接收并监听父进程传递的socket 描述符
// 4、子进程启动成功之后，父进程停止接收新连接， 同时等待旧连接处理完毕
// 5、父进程退出，热重启完成

// 关键点
// 1、监控信号，父进程获取文件描述符
// 		tl, ok := listener.(*net.TCPListener)
// 		f, err := tl.File()
// 2、发起一个请求给父进程
// 3、进程获取文件描述符，新建一个子进程监听该文件描述符
// 4、父进程 延迟关闭

var (
	server   *http.Server
	listener net.Listener = nil

	graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
	message  = flag.String("message", "hello world", "message to send")
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	x := (*[2]uintptr)(unsafe.Pointer(message))
	h := [3]uintptr{x[0], x[1], x[1]}
	w.Write(*(*[]byte)(unsafe.Pointer(&h)))
}

func main() {
	var err error

	// 解析参数
	flag.Parse()

	http.HandleFunc("/test", handler)
	server = &http.Server{Addr: ":3333"}

	if *graceful {
		// 子进程监听父进程传递的 socket 描述符
		log.Println("listening on the existing file descriptor 3")

		// 子进程的 0，1，2 是预留给标准输入，标准输出，错误输出，所以3 是传递的socket 描述符
		// 应放在子进程的 3
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		// 父进程监听 新建的socket 描述符
		log.Println("listening on a new file descriptor")
		listener, err = net.Listen("tcp", server.Addr)
	}

	if err != nil {
		log.Fatalf("listener error: %v", err)
	}

	go func() {
		err = server.Serve(listener)
		log.Printf("server.Serve err : %v \n", err)
	}()

	// 监听信号
	handleSignal()
	log.Println("signal end")

	//server.ListenAndServe()
}

func handleSignal() {
	c := make(chan os.Signal, 1)
	// 监听信号
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-c
		log.Printf("signal receive: %v \n", sig)
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Println("shutdown")
			signal.Stop(c)
			server.Shutdown(ctx)
			log.Println("graceful shutdown")
			return
		case syscall.SIGUSR2: // 进程热重启
			log.Println("reload")
			err := reload()
			if err != nil {
				log.Fatalf("graceful reload error: %v", err)
			}
			server.Shutdown(ctx)
			log.Println("graceful reload")
			return
		}
	}
}

func reload() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}

	// 获取socket 描述符
	f, err := tl.File()
	if err != nil {
		return err
	}

	// 设置传递给子进程的参数 （包含socket 描述符）
	args := []string{"-graceful"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout         // 标准输出
	cmd.Stderr = os.Stderr         // 错误输出
	cmd.ExtraFiles = []*os.File{f} //文件描述符

	//新建并执行子进程
	return cmd.Start()
}

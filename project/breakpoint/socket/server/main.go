package main

import (
	"io"
	"net"
)

/*
1、创建连接
2、发送上次接收到的文件内容位置
3、客户端就从上次断点的位置继续发送文件内容
4、客户端发送文件内容完毕，通知服务端
5、断开连接
*/

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func serverConn(conn net.Conn) {
	for {
		var buf = make([]byte, 10)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("server io EOF")
				return
			}
			log.Fatalf("server read faild: %s\n", err)
		}

		log.Printf("recevice %d bytes, content is 【%s】\n", n, string(buf[:n]))
	}
}

func main() {
	// 建立监听
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("error listen")
	}

	defer l.Close()

	log.Println("waiting accept.")

	var errChan = make(chan error)

	go func() {
		conn, err := l.Accept()
		defer conn.Close()
		if err != nil {
			log.Fatalf("accept faild: %s\n", err)
			errChan <- err
		}
		serverConn(conn)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	log.Fatalf("server out err: %d", <-errChan)
}

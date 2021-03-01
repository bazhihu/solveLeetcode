package main

import (
	"io"
	"net"
	"strconv"
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
	"log"
	"os"
	"os/signal"
	"syscall"
)

// 把接收到的内容 写入文件
func writeFile(content []byte) {
	if len(content) != 0 {
		fp, err := os.OpenFile("test_1.txt", (os.O_CREATE | os.O_WRONLY | os.O_APPEND), 0755)
		defer fp.Close()
		if err != nil {
			log.Fatal("open file faild: %s\n", err)
		}
		_, err = fp.Write(content)
		if err != nil {
			log.Fatal("append content to file faild: %s\n", err)
		}
		log.Printf("append content: 【%s】 success\n", string(content))
	}
}

// 获取已接受的内容的大小
// 断点续传需要把已接收内容大小 通知客户端从哪儿开始发送文件内容
func getFileStat() int64 {
	fileInfo, err := os.Stat("test_1.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("file size: %d\n", 0)
			return int64(0)
		}
		log.Fatalf("get file stat faild: %s\n", err)
	}
	log.Printf("file size: %d\n", fileInfo.Size())
	return fileInfo.Size()
}

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

		switch string(buf[:n]) {
		case "start-->":
			off := getFileStat()

			stringoff := strconv.FormatInt(off, 10)
			_, err := conn.Write([]byte(stringoff))
			if err != nil {
				log.Fatalf("server write faild: %s\n", err)
			}
			continue
		case "<--end":
			// 如果接收到客户端同志所有文件内容发送完毕消息则退出
			log.Fatalf("receive over \n")
			return
		}
		writeFile(buf[:n])
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

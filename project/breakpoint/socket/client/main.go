package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func clientRead(conn net.Conn) int {
	buf := make([]byte, 5)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("receive server info faild:%s \n", err)
	}

	// string conver int
	off, err := strconv.Atoi(string(buf[:n]))
	if err != nil {
		log.Fatalf("string conver int faild:%s\n", err)
	}
	return off
}

func clientWrite(conn net.Conn, data []byte) {
	_, err := conn.Write(data)
	if err != nil {
		log.Fatalf("send 【%s】 content faild: %s\n", string(data), err)
	}
	log.Printf("send 【%s】 content success\n", string(data))
}

func clientConn(conn net.Conn) {
	defer conn.Close()

	// 发送 "start-->"消息通知服务端，我要开始发送文件内容了
	// 获取服务端的通知，得知已经接收到了多少内容，从已经接收的内容处开始继续发送
	clientWrite(conn, []byte("start-->"))
	off := clientRead(conn)

	// send file content
	fp, err := os.OpenFile("test.txt", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatalf("open file faild: %s\n", err)
	}
	defer fp.Close()

	// 设置从那儿开始读取文件
	_, err = fp.Seek(int64(off), 0)
	if err != nil {
		log.Fatalf("set file seek faild: %s\n", err)
	}

	log.Printf("read file at seek: %d\n", off)

	// start send
	for {
		// 每次发送10个字节大小的内容
		data := make([]byte, 10)
		n, err := fp.Read(data)
		if err != nil {
			if err == io.EOF {
				// 如果已经读取完 文件内容
				// 就发送 '<--end' 消息通知服务端， 文件内容发送完了
				time.Sleep(time.Second * 1)
				clientWrite(conn, []byte("<-- end"))
				log.Println("send all content, now quit")
				break
			}
			log.Fatalf("read file err: %s\n", err)
		}
		// 发送文件内容到服务端
		clientWrite(conn, data[:n])
	}
}

func main() {
	// connect timeout 10s
	conn, err := net.DialTimeout("tcp", ":8888", time.Second*10)
	if err != nil {
		log.Fatalf("client dial faild: %s\n", err)
	}
	clientConn(conn)
}

package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// 造成粘包原因
// 1、由Nagle算法造成得发送端的粘包：Nagle算法是一种改善网络传输效率的算法。
// 简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，
// 而是等待一小段时间看看再等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。

// 接受端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层数据。
// 当应用层由于某些原因不能及时把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。

// 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	//读取消息长度
	lenghtByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lenghtBuff := bytes.NewBuffer(lenghtByte)
	var lenght int32
	err := binary.Read(lenghtBuff, binary.LittleEndian, &lenght)
	if err != nil {
		return "", err
	}

	// 判断buffered返回缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < lenght+4 {
		return "", err
	}
	// 读取真正的消息数据
	pack := make([]byte, int(4+lenght))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

// 监听客户端接受，处理接受客户端信息数据
func process(conn net.Conn) {
	defer conn.Close()
	// 从请求连接中，获取bufio 流
	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decide msg failed, err:", err)
			return
		}
		fmt.Println("收到client 发来的数据：", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()

	// 接收请求, accept返回的一个请求连接conn
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

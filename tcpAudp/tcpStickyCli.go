package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// 造成粘包原因
// 1、由Nagle算法造成得发送端的粘包：Nagle算法是一种改善网络传输效率的算法。
// 简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，
// 而是等待一小段时间看看再等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。

// 接受端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层数据。
// 当应用层由于某些原因不能及时把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。

// 消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息长度
	var length int32 = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `hello, hello.How are you?`
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err：", err)
			return
		}
		conn.Write(data)
	}
}

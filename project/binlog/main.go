package main

import (
	"fmt"
	"net"
	"os"
	"solveLeetcode/project/binlog/packet"
	"time"
)

func main() {
	// 连接mysql
	conn, err := connMysql("127.0.0.1:3306", "root", "password")
	if err != nil {
		fmt.Println("connMysql-err", err)
		os.Exit(1)
	}
	// 进行握手 HandShake协议
	handshake(conn)
	// show global variables like 'binlog_checksum' 返回checksum 类型

	// show master status 返回binlogFileName, Position

	// 发送 binlog dump 命令

	// 循环读取binlog event 流
}

func connMysql(addr string, user string, password string) (*packet.Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	return packet.NewConn(conn), err
}

func handshake(conn *packet.Conn) error {

	data, err := conn.ReadPacket()
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("data", data)
	return nil
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

var (
	QUERY = uint8(3)
)

func main() {
	var errChan = make(chan error)
	conn, err := net.DialTimeout("tcp", "127.0.0.1:3306", 10*time.Second)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(-1)
	}

	fmt.Println("conn-success", conn)

	go read(conn)

	send(conn, `SHOW GLOBAL VARIABLES LIKE 'BINLOG_CHECKSUM'`)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	error := <-errChan

	fmt.Println("err", error)
}

// read mysql data
func read(conn net.Conn) {
	go processBinlog(conn)
}

func processBinlog(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
	}
}

func send(conn net.Conn, command string) {
	arg := makeHead(getCommandBuf([]byte(command)))

	if n, err := conn.Write(arg); err != nil || n != len(arg) {
		fmt.Println("err", err)
		return
	}
	return
}

func makeHead(arg []byte) []byte {
	length := len(arg) - 4
	arg[0] = byte(length)
	arg[1] = byte(length >> 8)
	arg[2] = byte(length >> 16)
	arg[3] = uint8(0)
	return arg
}

// header has 4 bytes
func getCommandBuf(arg []byte) []byte {
	length := len(arg) + 1 + 4
	// new array byte
	data := make([]byte, length)

	data[4] = QUERY
	copy(data[5:], arg)
	return data
}

func StringToByteSlice(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

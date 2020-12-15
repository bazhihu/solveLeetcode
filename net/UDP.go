package main

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:303")
	lis.Accept()
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		fmt.Println("data", string(data[:n]), addr, n)

		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			continue
		}
	}
}

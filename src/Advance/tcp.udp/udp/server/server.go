package main

import (
	"fmt"
	"net"
)

// UDP Server端

func main() {
	// net.Listen("udp","127.0.0.1:55555")
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("Listen UDP failed, err: ", err)
		return
	}
	defer conn.Close()
	// 此处不需要再建立连接了，直接收发数据即可
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("ReadFromUDP failed, err: ", err)
			return
		}
		fmt.Println(string(data[:n]))
		msg := "Received Success!"
		conn.WriteToUDP([]byte(msg), addr)
	}
}

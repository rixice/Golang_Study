package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP Client端

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	defer conn.Close()
	if err != nil {
		fmt.Println("连接服务器失败, err: ", err)
	}
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入：")
		msg, _ := reader.ReadString('\n')
		conn.Write([]byte(msg))
		// 接收回复的消息
		n, _, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("Receive reply msg failed, err: ", err)
			return
		}
		fmt.Println("收到回复：", string(reply[:n]))
	}

}

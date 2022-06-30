package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server端

func goProcess(conn net.Conn) {
	// 3. 与客户端通信
	var tmp [128]byte
	var response [256]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("Read from conn failed, err: %v\n", err)
			return
		}
		fmt.Printf(string(tmp[:n]))
		fmt.Printf("请回复: ")
		fmt.Scanln(&response)
		input, _ := reader.ReadString('\n')
		msg := strings.TrimSpace(input)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

}

func main() {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Start listen failed, err: %v\n", err)
		return
	}
	// 2. 等待客户端建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept connection failed, err: %v\n", err)
			return
		}
		go goProcess(conn)
	}
}

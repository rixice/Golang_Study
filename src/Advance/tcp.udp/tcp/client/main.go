package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client端

func main() {
	// // 1. 与server端建立连接
	// conn, err := net.Dial("tcp", "127.0.0.1:20000")
	// if err != nil {
	// 	fmt.Printf("dial 127.0.0.1:20000 failed, err: %v\n", err)
	// 	return
	// }
	// // 2. 发送数据
	// var msg string
	// var arg string
	// for {
	// 	fmt.Scanln(&arg)
	// 	if arg == "0" {
	// 		conn.Close()
	// 		return
	// 	}
	// 	msg = arg
	// 	conn.Write([]byte(msg))
	// }
	/////////////////////////////////////////////
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial 127.0.0.1:20000 failed, err: %v\n", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin) // 读取用户输入
	for {
		input, _ := inputReader.ReadString('\n') // 识别到换行(回车)才结束用户输入
		inputInfo := strings.Trim(input, "\r\n") // 将windows自带的换行符去掉
		if strings.ToUpper(inputInfo) == "Q" {   // 如果输入q或Q，就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{} // 读取服务端返回的消息才能继续发送消息
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("recv failed, err: %v\n", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

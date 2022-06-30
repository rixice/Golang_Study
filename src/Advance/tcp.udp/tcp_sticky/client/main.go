package main

import (
	"fmt"
	proto "golang/src/go_code/Internet/day01/tcp_sticky/protocol"
	"net"
)

// 黏包 client

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("Dial failed, err: ", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintln("How are you?")
		// 调用协议编码数据
		b, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("Encode failed, err: ", err)
			return
		}
		conn.Write(b)
		// time.Sleep(time.Second)
	}
}

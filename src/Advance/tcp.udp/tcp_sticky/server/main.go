package main

import (
	"bufio"
	"fmt"
	proto "golang/src/go_code/Internet/day01/tcp_sticky/protocol"
	"io"
	"net"
)

// 引申知识点：
// 大端Big endian：多个字节的数，将高位写在内存的左端<低地址>（从左往右读）
// 小端Little endian：将高位写在内存的右端<高地址>（由右往左读）

func goProcess(conn net.Conn) {
	// 3. 与客户端通信
	defer conn.Close()
	// var buf [1024]byte
	reader := bufio.NewReader(conn)
	for {
		// n, err := reader.Read(buf[:])
		// if err == io.EOF {
		// 	break
		// }
		// if err != nil {
		// 	fmt.Println("read from client failed, err: ", err)
		// 	break
		// }
		// recvStr := string(buf[:n])
		recvMsg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Decode failed, err: ", err)
			return
		}
		fmt.Println("收到client发来的数据：", recvMsg)
		// 结果20条消息，实际只分了三次发送	<黏包>
		// 在客户端加了sleep 1秒后，就变成了20次发送，但这样影响效率

		// 如何解决黏包：
		// 对数据包进行封包和拆包，给一段数据加上包头，包头长度固定并存储了包体的长度
	}

}

func main() {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("Start listen failed, err: %v\n", err)
		return
	}
	defer listener.Close()
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

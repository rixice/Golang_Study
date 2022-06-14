package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件写入操作

func main() {
	fileWrite_2()
}

func fileWrite_1() {
	fileObj, err := os.OpenFile("C:\\Users\\zhang_jianxiong\\Desktop\\test.txt", os.O_APPEND|os.O_CREATE, 777)
	// windows下文件权限随意，不生效
	if err != nil {
		fmt.Printf("File Open Failed, err: %v\n", err)
		return
	}
	// write
	fileObj.Write([]byte("Hello World哈哈哈\n"))
	// writeString
	fileObj.WriteString("")
	fileObj.Close()
}

func fileWrite_2() {
	fileObj, err := os.OpenFile("C:\\Users\\zhang_jianxiong\\Desktop\\test.txt", os.O_APPEND|os.O_CREATE, 777)
	// windows下文件权限随意，不生效
	if err != nil {
		fmt.Printf("File Open Failed, err: %v\n", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("HELLO\n") // 写到缓存中
	wr.Flush()                // 将缓存中的内容写入文件
}

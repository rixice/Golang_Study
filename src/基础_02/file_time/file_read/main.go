// 文件操作
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件
func main() {
	readFromFileIoutil()
}

func readFromFile_1() {
	fileObj, err := os.Open("C:\\Users\\zhang_jianxiong\\Desktop\\test.txt") //传一个文件的路径
	if err != nil {
		fmt.Printf("open file failed, error: %v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	var tmp = make([]byte, 128) // 指定读的长度
	for {
		n, err := fileObj.Read(tmp)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Read file failed, error: %v\n", err)
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// 利用bufio这个包读取文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("C:\\Users\\zhang_jianxiong\\Desktop\\test.txt") //传一个文件的路径
	if err != nil {
		fmt.Printf("open file failed, error: %v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 创建一个用来从文件中读取的对象
	for {
		reader := bufio.NewReader(fileObj)
		line, err := reader.ReadString('\n') // 指定字符
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err: %v", err)
			return
		}
		fmt.Print(line)
	}
}

// 利用ioutil读文件（简单）
func readFromFileIoutil() {
	ret, err := ioutil.ReadFile("C:\\Users\\zhang_jianxiong\\Desktop\\test.txt") //传一个文件的路径
	if err != nil {
		fmt.Printf("open file failed, error: %v\n", err)
		return
	}
	// ioutil自动关闭文件
	fmt.Println(string(ret))
}

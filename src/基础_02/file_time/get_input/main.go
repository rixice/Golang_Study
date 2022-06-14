package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时，如果有空格
func main() {
	str := make(map[string]interface{}, 8)
	str["haha"] = 1
	str["pupu"] = "hello"
	fmt.Println(str)
	for _, v := range str {
		fmt.Println(v)
	}
}

func useBufio(s string) {
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
}

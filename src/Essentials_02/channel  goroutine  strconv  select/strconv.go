package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	str := "10000"
	// 参数：字符串、进制数、位数
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}
	fmt.Printf("%T\n", ret) // 默认返回int64

	// 如果只想将string转为int
	var ret_int int
	ret_int, _ = strconv.Atoi(str)
	fmt.Printf("%T\n", ret_int)

	// int转为string
	var ret_str string
	i := 999
	ret_str = strconv.Itoa(i)
	fmt.Printf("%T\n", ret_str)
}

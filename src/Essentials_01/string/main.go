// 单独的字母、汉字、符号表示一个字符

// 字节：1字节=8Bit（8个二进制位）
// 1个字符'A' = 1个字节
// 1个utf8编码的汉字 = 一般占3个字节

// 注意：Go中，字符串是用双引号包裹的  "你好"
//            单引号包裹的是字符   '你'

package main

import (
	"fmt"
	"strings"
)

func main() {
	// \ 具有转义功能
	path := "\"D:\\Go\\src\\code\""
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	// 多行的字符串
	// ``反引号的内容呈原样输出
	// s2 := `
	// 111
	// 	222
	// 		333
	// `
	s3 := `"D:\Go\src\code"`
	// fmt.Println(s2, s3)

	// 字符串拼接
	name := "张三"
	word := "dsb"
	ss := name + word
	ss1 := fmt.Sprintf("%s%s", name, word)
	fmt.Println(ss, ss1)

	// 分割
	ret := strings.Split(s3, "\\") // 相当于弄成了数组
	fmt.Println(ret)
	// 判断包含某个字符
	fmt.Println(strings.Contains(ss, "张三"))
	fmt.Println(strings.Contains(ss, "张四"))
	// 前缀
	fmt.Println(strings.HasPrefix(ss, "张"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "张"))

	// 检查指定字符的第一次出现的位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "b"))
	// 检查指定字符的最后一次出现的位置
	fmt.Println(strings.LastIndex(s4, "b"))

	// 拼接
	fmt.Println(strings.Join(ret, "+")) // 需要是分割后的字符串
}

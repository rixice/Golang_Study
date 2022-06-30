package main

import (
	"fmt"
)

// 浮点数
func main() {
	//math.MaxFloat32 // float32最大值
	f1 := 1.23456
	fmt.Printf("%T", f1) // 默认Go中的小数都是float64类型
	// float32的变量不能直接赋值给float64
}

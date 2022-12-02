package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller()

func main() {
	f1()
}

func f1() {
	// Caller里面的数字代表往上找多少层（函数被多次调用时）
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(path.Base(file))
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A: %d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B: %d\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)         // 设置只占用1个逻辑CPU <线程>（可以超过实际CPU核心数）
	fmt.Println(runtime.NumCPU()) // 打印总的线程数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

// 开发人员必须要确认goroutine在哪里结束 以及 goroutine的调度模型(GMP)
// goroutine初始栈的大小——2KB

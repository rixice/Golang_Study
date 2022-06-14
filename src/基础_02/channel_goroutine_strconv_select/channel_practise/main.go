package main

import (
	"fmt"
	"sync"
)

// channel练习
// 1. 启动一个goroutine，生成100个数发送到ch1中
// 2. 启动一个goroutine, 从ch1中取值，计算其平方放到ch2中
// 3. 在main中 从ch2取值打印出来

var wg sync.WaitGroup

func func1(ch1 chan int) {
	// 单向通道
	// ch1 chan<- int(只能往里面发送值)
	// <-ch1 单向发送通道，不能用来取值，会报错
	// ch1 <-chan int(只能往里面取出值)
	// ch1 <- n 单项取值通道，往里面发送值，会报错
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) // 不关闭通道无法进行遍历！！！
}

func func2(ch1, ch2 chan int) {
	defer wg.Done()
	for i := range ch1 {
		ch2 <- i * i
	}
	close(ch2) // 不关闭通道无法进行遍历！！！
}

// 如果goroutine结束了还持有chan，那么chan就不会被释放
func main() {
	wg.Add(2)
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go func1(ch1)
	go func2(ch1, ch2)
	wg.Wait()
	for ret := range ch2 {
		fmt.Println(ret)
	}
}

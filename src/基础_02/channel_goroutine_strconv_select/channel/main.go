package main

import (
	"fmt"
	"sync"
)

// 虽然可以使用共享内存进行数据交换，但这样会存在不同的goroutine中发生竞态问题
// 所以必须使用互斥量对内存进行加锁，这种做法势必造成性能问题
// Go中的并发模型是CSP，通过通信共享内存，而不是通过共享内存实现通信

// channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制

var a []int
var c chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	c = make(chan int) // <不带缓冲区的> 通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-c
		fmt.Println("后台goroutine从通道c中取到了", x)
	}()
	c <- 10 // 卡住了 <死锁>
	fmt.Println("10发送到通道b中了。。。。")
	wg.Wait()
}

func bufChannel() {
	c = make(chan int, 16) // <带缓冲区的> 通道的初始化
	c <- 10                // 不会卡住
	x := <-c
	fmt.Println(x) // 10
	close(c)       // 最好是执行完关闭（但也会自动回收）
}

func main() {
	bufChannel()
}

// 通道的操作：
// 1. 发送： ch <- 1
// 2. 接收： x := <- ch
// 3. 关闭： close()
// 通道的缓冲区尽可能小

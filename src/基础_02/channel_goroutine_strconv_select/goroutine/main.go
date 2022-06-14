package main

import (
	"fmt"
	"time"
)

//////////////////////////////////////////////////////////
// goroutine的调度
// GMP是Go语言运行时（runtime）层面的实现，是go自己实现的一套调度系统。区别于操作系统调度OS线程

// G就是goroutine，里面除了由本goroutine信息外，还有与所在P的绑定等信息

// M是Go运行时对操作系统内核线程的虚拟，M与内核线程一般是一对一映射的关系，一个goroutine最终要放到M上执行

// P管理着一组goroutine队列，里面会存储当前goroutine运行的上下文环境;
// P会对自己管理的goroutine队列做一些调度，当自己的队列消费完了就去全局队列里取;
// 如果全局队列也消费完了，就会去其他P里面取
//////////////////////////////////////////////////////////
// func hello() {
// 	fmt.Println("hello world")
// }
// 程序启动之后会创建一个主goroutine去执行
func main() {
	// go hello()              // 开启一个单独的goroutine去执行hello函数（任务）
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println(i) // 用的是函数参数的那个i，不是外面的i
		}(i)
	}
	time.Sleep(time.Second) // main函数结束的话，由main函数产生的goroutine也都结束了
	fmt.Println("main")
}

// goroutine什么时候结束
// goroutine对应的函数结束了，goroutine就结束了。

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

// func f() {
// 	// UnixNano()代表1970年到现在的纳秒数，保证Seed不一样
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 5; i++ {
// 		r1 := rand.Int()    // int64的随机数
// 		r2 := rand.Intn(10) // 返回一个 0~100 的数
// 		fmt.Println(0-r1, 0-r2)
// 		f1(i)
// 	}
// }

func f1(i int) {
	defer wg.Done() // 每执行一次，计数器减1
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	// fmt.Println(time.Now().UnixNano())
	// f()
	// wg.Add(10)	循环外也可以写
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	// 如何知道这10个goroutine都结束了
	// 尽量别用time.sleep()
	wg.Wait() // 等待wg的计数器减为0
}

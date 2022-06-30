package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	// lock.Lock()
	// x++
	atomic.AddInt64(&x, 1) // 使用内置的包进行这个加法操作
	// lock.Unlock()
	wg.Done()
}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {

		go add()

	}
	wg.Wait()
	fmt.Println(x)
}

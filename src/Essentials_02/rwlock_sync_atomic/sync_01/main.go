package main

import (
	"fmt"
	"sync"
)

// 锁

// 但凡涉及到公共资源的访问，都要加锁！
// 如果是 读远远大于写 的场景，互斥锁就会非常影响性能

// 如果某个操作只想执行一次，就用sync.Once()!!!
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	lock.Lock() // 尽量写for外面，避免串行任务
	for i := 0; i < 500000; i++ {
		x += 1
	}
	lock.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

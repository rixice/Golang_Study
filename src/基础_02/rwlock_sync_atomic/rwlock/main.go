package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock 读写互斥锁

var (
	x      = 0
	lock   sync.Mutex
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

// 读操作：加的是读锁
func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Microsecond)
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	x += 1
	time.Sleep(time.Microsecond * 30)
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

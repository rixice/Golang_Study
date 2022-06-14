package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go内置的map不是并发安全的

// var m = make(map[string]int)
// var lock sync.Mutex

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int) {

// 	m[key] = value

// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			lock.Lock()
// 			set(key, n)
// 			lock.Unlock()
// 			fmt.Printf("key=%v\tvalue=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

var m2 = sync.Map{} // Go提供sync.Map来进行Map的并发安全

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 必须使用sync.Map内置的Store方法存值
			value, _ := m2.Load(key) // 必须使用sync.Map提供的Load方法取值
			fmt.Printf("key=%v\tvalue=%v\n", key, value)
			wg.Done()
		}(i) // 匿名函数传参
	}
	wg.Wait()
}

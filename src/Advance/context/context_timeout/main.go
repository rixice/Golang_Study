package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("MySQL Connecting...")
		time.Sleep(time.Second * 1) // 假设正常连接数据库需要1秒
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("Worker Done!")
	wg.Done()
}

func main() {
	// 设置一个超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知 子goroutine 结束
	wg.Wait()
	fmt.Println("Over!")
}

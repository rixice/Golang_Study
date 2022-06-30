package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context？
// 当一个goroutine内又再次启动了其他的goroutine(e.g: http Server中的众多服务)
// 当 根goroutine的ctx 超时了,那么其下的所有goroutine都将超时

// Context 是 Go v1.7 之后加入的标准库,用于简化对处理单个请求的多个goroutine之间与
// 请求域的数据,取消信号,截止时间等相关操作

var wg sync.WaitGroup
var notify bool
var notify_ch = make(chan bool, 1)

func f1(ctx context.Context) {
	defer wg.Done()
LOOP: // 指定标签
	for {
		fmt.Println("Hahahaha")
		time.Sleep(time.Second)
		// if notify {
		// 	break
		// }
		select {
		case <-ctx.Done():
			break LOOP // 跳出指定标签
		default:
		}
	}
}

// 如果用context来解决这个超时问题,怎么解决?
func f(ctx context.Context) {
	defer wg.Done()
LOOP: // 指定标签
	for {
		fmt.Println("Hello world")
		time.Sleep(time.Second)
		// if notify {
		// 	break
		// }
		go f1(ctx)
		select {
		case <-ctx.Done():
			break LOOP // 跳出指定标签
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 3)
	// notify = true
	cancel()
	wg.Wait()
	// 如何通知 子goroutine 退出
}

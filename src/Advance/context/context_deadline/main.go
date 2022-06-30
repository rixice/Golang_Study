package main

import (
	"context"
	"fmt"
	"time"
)

// context.WithDeadline

func main() {
	d := time.Now().Add(5000 * time.Millisecond) // 这里设置了超时时间
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel()都是很好的实践
	// 如果不这样做，可能会使上下文及其父类存活时间超过其必要的时间
	defer cancel()

	select {
	case <-time.After(1 * time.Second): // 1秒后才输出
		fmt.Println("Hello")
	case <-ctx.Done(): // 超时则输出以下err
		fmt.Println(ctx.Err())
	}
}

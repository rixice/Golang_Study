package main

import (
	"context"
	"fmt"
)

func getNum(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() { // return结束该goroutine，防止泄露
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range getNum(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

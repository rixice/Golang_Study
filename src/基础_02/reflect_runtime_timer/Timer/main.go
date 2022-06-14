package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t) // 1秒执行一次
	}
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序
1. 开启一个goroutine循环生成int64类型的随机数，发送的jobChan
2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3. 主goroutine从resultChan取出结果并打印到终端输出
*/

type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

var wg sync.WaitGroup
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func func1(c1 chan<- *job) {
	defer wg.Done()
	go func() {
		for {
			x := rand.Int63()
			newJob := &job{
				value: x,
			}
			c1 <- newJob
			time.Sleep(time.Second * 1)
		}
	}()
}

func func2(c1 <-chan *job, c2 <-chan *result) {
	defer wg.Done()
	go func() {
		for {
			x := <-c1
			sum := int64(0)
			n := x.value
			for n > 0 {
				sum += n % 10
				n = n / 10
			}
			newResult := &result{
				job:    x,
				result: sum,
			}
			resultChan <- newResult
		}
	}()
}

func main() {
	// n := 123
	// for n > 0 {
	// 	fmt.Println(n % 10) // 3 2 1	(取各位数)
	// 	n = n / 10          // 12 1 0
	// }
	wg.Add(1)
	go func1(jobChan)
	// 开启24个goroutine
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go func2(jobChan, resultChan)
	}
	for num := range resultChan {
		fmt.Printf("value: %v\tsum：%v\n", num.job.value, num.result)
	}
	wg.Wait()
}

package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	<-ch1
	<-ch1
	x, ok := <-ch1
	if !ok {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(x)
}

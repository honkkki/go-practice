package main

import "fmt"

func go1(ch chan int) {
	fmt.Println("go")
	<-ch
}

func main() {
	var ch = make(chan int)		// 无缓冲区的管道
	go go1(ch)
	ch <- 1
	fmt.Println("main")
}

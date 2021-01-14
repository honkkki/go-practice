package main

import "fmt"

func main()  {
	ch := make(chan int, 6)		// 有缓冲区的管道
	//<-ch		// 死锁
	ch <- 1
	ch <- 1
	fmt.Println("hello")

}
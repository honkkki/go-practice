package main

import "fmt"

func go1(ch chan int) {
	<-ch
	fmt.Println("go")
}

func main() {
	var ch = make(chan int)
	go go1(ch)
	ch <- 1
	fmt.Println("main")
}

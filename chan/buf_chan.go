package main

import "fmt"

func test(ch chan int, ch2 chan int) {
	for v := range ch {
		fmt.Println(v)
		ch2 <- v
	}

	close(ch2)
}

func main() {
	ch := make(chan int, 6)
	ch2 := make(chan int, 6)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	go test(ch, ch2)
	fmt.Println("main")

	for _ = range ch2 {
	}
}

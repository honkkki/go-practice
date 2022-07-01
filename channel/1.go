package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 1

	for {
		select {
		case d := <-ch:
			fmt.Println(d)

		default:
			fmt.Println("finish")
			return
		}
	}

}

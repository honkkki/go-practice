package main

import "fmt"

func main() {
	ch := make(chan int, 6)

	ch <- 1
	ch <- 1
	ch <- 1

	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

	//_, ok := <- ch
	//fmt.Println(ok)

}

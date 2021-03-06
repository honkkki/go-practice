package main

import (
	"fmt"
)

func main() {
	var ch1 = make(chan int, 5)
	fmt.Println(ch1)
	fmt.Printf("%T \n", ch1)

	ch1 <- 10
	ch1 <- 2
	x := <-ch1
	y := <-ch1

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("-------------")

	ch1 <- 1
	ch1 <- 1
	ch1 <- 1
	ch1 <- 1
	close(ch1)
	c := <-ch1
	fmt.Println(c)
	fmt.Println("-------------")

	for v := range ch1 {
		fmt.Println(v)
	}

	cc := <-ch1
	fmt.Println(cc)

}

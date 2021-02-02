package main

import "fmt"


func main() {
	var ch1 = make(chan int, 5)
	fmt.Println(ch1)
	fmt.Printf("%T \n", ch1)

	ch1 <- 10
	ch1 <- 2
	x := <-ch1
	fmt.Println(111)
	y := <-ch1

	fmt.Println(x)
	fmt.Println(y)
}

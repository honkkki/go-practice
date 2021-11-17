package main

import (
	"fmt"
	"unsafe"
)

func main() {
	c := make(chan struct{}, 10)
	c<- struct{}{}
	s := <-c
	fmt.Println(unsafe.Sizeof(s))

	c1 := make(chan int)
	//c2 := make(chan int)
	//m := make(map[chan int]struct{})
	fmt.Println(c1 == c1)


}

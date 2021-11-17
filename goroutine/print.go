package main

import (
	"fmt"
)

// 交替打印数组
func main() {
	s1 := []int{0, 2, 4, 6, 8}
	s2 := []int{1, 3, 5, 7, 9}

	ch := make(chan struct{})

	go func() {
		for _, v := range s1 {
			<-ch
			fmt.Println(v)
			ch <- struct{}{}
		}
	}()

	for _, v := range s2 {
		ch <- struct{}{}
		fmt.Println(v)
		<-ch
	}

}

package main

import (
	"fmt"
)

func main() {
	// 二维切片
	s := make([][]int, 10)
	s[0] = make([]int, 1)
	s[0][0] = 1
	fmt.Println(s)

	s1 := [][]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(s1)

	fmt.Println("--------------------------------")

	sli1 := []int{1, 2, 3}
	sli := make([]int, 0, 10)
	sli = append(sli, sli1...)
	fmt.Println(sli)
	sli2 := sli1[1:]
	fmt.Println(sli2)

}

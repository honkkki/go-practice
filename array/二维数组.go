package main

import "fmt"

func main() {
	// 二维数组
	a := [3][2]int{
		{1, 2},
		{1, 2},
		{1, 2},
	}

	fmt.Println(a)

	for k, val := range a {
		fmt.Printf("row[%d] = %v \n", k, val)
		fmt.Println()
	}

	s := [][]int{
		{1, 2},
		{1, 2},
		{1, 2},
	}

	fmt.Println(s)

}

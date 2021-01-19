package main

import "fmt"

func main() {
	arr1 := [2]string{"jisoo", "rose"}
	arr2 := [...]string{"jisoo", "rose"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Printf("%T \n", arr1)

	arr3 := [3]interface{}{
		1, "a", true,
	}

	fmt.Println(arr3)
	fmt.Printf("%T \n", arr3)

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

}

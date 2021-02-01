package main

import "fmt"

func main()  {
	// 切片是对底层数组的引用
	arr := [5]int{1,2,3,4,5}
	s := arr[:]
	arr[0] = 100

	fmt.Println(s)

	// 修改切片来修改数组
	s[1] = 200
	fmt.Println(s)
	fmt.Println(arr)



}

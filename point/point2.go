package main

import "fmt"

func main() {
	a := 1
	p := &a

	fmt.Println(p)

	pp := &p
	fmt.Println(pp)

	var p2 *int
	fmt.Println(p2) // nil
	p2 = &a
	*p2 = 100
	fmt.Println(*p2)
	fmt.Println(a)

	fmt.Println("-----------------")

	// 数组指针取值
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	arrPoint := &arr
	(*arrPoint)[0] = 100
	fmt.Println(arr)

	var s []int
	fmt.Println(s)

}

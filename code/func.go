package main

import "fmt"

type MyFunc func(num int) int

func CalculateNum(num int, myFunc2 MyFunc) int {
	res := myFunc2(num)
	return res
}

func main() {
	res := CalculateNum(1, func(num int) int {
		num++
		return num
	})

	res2 := CalculateNum(1, func(num int) int {
		num--
		return num
	})

	fmt.Println(res)
	fmt.Println(res2)
}

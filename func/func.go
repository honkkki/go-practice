package main

import "fmt"

func intSum(num ...int) int {
	res := 0
	for _, v := range num {
		res = res + v
	}

	return res
}

func testFunc() {
	fmt.Println(1)
}

func main() {
	res := intSum(10, 20, 30)
	fmt.Println(res)
	func1 := testFunc
	func1()
}

package main

import "fmt"

func ff() []int {
	a := []int{1, 2, 3}
	return a
}

func main() {
	s := ff()
	fmt.Println(&s)
	fmt.Println(ff()[1])
	//fmt.Println(&ff())	// cannot take the address 字面量无法取地址
	//fmt.Println([3]int{1,2,3}[:])		// error

}

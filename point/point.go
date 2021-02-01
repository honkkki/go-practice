package main

import "fmt"

func test(sli []int) {
	sli[0] = 2
}

func test2(num *int) {
	*num = 2
}

func main() {
	sli := []int{1, 2, 3}
	num := 1
	fmt.Println(sli)
	test(sli)
	fmt.Println(sli)

	fmt.Println(num)
	test2(&num)
	fmt.Println(num)

	a := new(int)
	fmt.Println(a) // 0xc00006e0e8
	fmt.Printf("%T", a)
}

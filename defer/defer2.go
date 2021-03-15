package main

import "fmt"

func main() {
	//i := 0
	//defer fmt.Println(i) // i=0 已经赋值
	//i = 1000
	//fmt.Println(i)

	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
		//defer fmt.Println(i)
	}
	//4
	//3
	//2
	//1
	//0

	defer fmt.Println("--------------")

	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
		//defer fmt.Println(i)
	}
	//5
	//5
	//5
	//5
	//5

}

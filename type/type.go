package main

import "fmt"

// 自定义类型
type MyInt int

// 类型别名
type NewInt = int

func main()  {
	var num MyInt
	fmt.Printf("%T \n", num)

	var num2 NewInt
	fmt.Printf("%T \n", num2)

	//main.MyInt
	//int
}

package main

import (
	"fmt"
)

var name = "name"

func test()  {
	name = "jisoo"
	var bytes byte
	fmt.Println(name)
	fmt.Println(bytes)
	fmt.Println("a\na")
}

func main()  {
	test()
	fmt.Println("goland")
	var age int
	var num float32
	fmt.Println(age, num, name)
}




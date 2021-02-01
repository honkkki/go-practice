package main

import (
	"fmt"
	"strconv"
)

var name = "name"

func test()  {
	name = "jisoo"
	var bytes byte
	fmt.Println(name)
	fmt.Println(bytes)
}

func main()  {
	test()
	fmt.Println("---------------------------------")
	fmt.Println("goland")
	var age int
	var num float32
	fmt.Println(age, num, name)


	s := "10"
	i, _ := strconv.Atoi(s)
	fmt.Println(i)

	p := new(int)
	fmt.Println(p)





}




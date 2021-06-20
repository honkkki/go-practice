package main

import (
	"fmt"
	"reflect"
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

	fmt.Println("---------------------------------")

	str := "hello"
	strBytes := []rune(str)
	fmt.Println(strBytes[0])
	fmt.Println(string(strBytes[0] - 32))
	strBytes[0] = 'H'
	fmt.Println(reflect.TypeOf('H'))		// rune
	str = string(strBytes)
	fmt.Println(str)



}




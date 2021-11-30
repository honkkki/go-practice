package main

import "fmt"

type student struct {
	name string
}

type person struct {
	name string
}

func (s student) read() {
	fmt.Println("read")
}

func main() {
	s := student{name: "gopher"}
	p := person{name: "gopher"}
	// 两个属性完全相同的结构体可以互相转换
	s = student(p)
	fmt.Println(s)
	fmt.Println(p)

}

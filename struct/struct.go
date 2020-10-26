package main

import "fmt"

type Person struct {
	name string
	age int
}

func main()  {
	person := Person{
		name: "jisoo",
		age:  20,
	}

	fmt.Println(person)

	var st Person
	fmt.Println(st)

	// 结构体指针
	p1 := &Person{}
	fmt.Println(p1)
	p1.name = "jisoo"
	p1.age = 20

	var ptr = new(Person)
	fmt.Println(ptr)
	fmt.Printf("%#v \n", p1)

	// 匿名结构体
	var st1 struct{
		name string
	}

	st1.name = "jisoo"



}

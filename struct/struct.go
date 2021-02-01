package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{
		name: "jisoo",
		age:  20,
	}

	fmt.Println(person)

	fmt.Println("----------------")

	var st Person
	fmt.Println(st)

	var p *Person
	fmt.Println(p) // nil
	//p.name = "hello"	// panic
	//fmt.Println(p)

	// 结构体指针
	p1 := &Person{}
	fmt.Println(p1)
	p1.name = "jisoo"
	p1.age = 20
	fmt.Println(p1)

	var ptr = new(Person)
	fmt.Println(ptr)
	fmt.Printf("%#v \n", p1)

}

package main

import "fmt"

// 对外部包可见
type Person struct {
	name string
	age  int
}

type integer int

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func (i integer) Print() {
	fmt.Println(i)
}

func (i *integer) pt() {
	*i = 1
	fmt.Println("val:", *i)
}

func main() {
	p := NewPerson("jisoo", 18)
	fmt.Println(p)
	p.SetAge(20)
	fmt.Println(p)
	fmt.Println("---------------")

	var i integer
	i = 100
	i.Print()
	fmt.Println("---------------")

	var ii integer
	fmt.Println("ii:", ii)
	ii.pt()
	fmt.Println("ii:", ii)
	fmt.Println("---------------")

	iii := new(integer)
	fmt.Println(iii)
	iii.pt()
}

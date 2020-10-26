package main

import "fmt"

// 对外部包可见
type Person struct {
	name string
	age int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) SetAge(newAge int)  {
	p.age = newAge
}

func main()  {
	p := NewPerson("jisoo", 18)
	fmt.Println(p)
	p.SetAge(20)
	fmt.Println(p)

}

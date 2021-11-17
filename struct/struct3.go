package main

import "fmt"

// Animal 结构体继承
type Animal struct {
	Name string
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) move()  {
	fmt.Println("dog move")
}

func (a *Animal) move()  {
	fmt.Printf("%s can move \n", a.Name)
}

func (d *Dog) wang()  {
	fmt.Printf("%s can wangwangwang \n", d.Name)
}

func main()  {
	d := Dog{
		Feet:   4,
		Animal: &Animal{
			Name: "dog",
		},
	}

	d.Animal.move()
	d.move()
	d.wang()
	fmt.Println(d.Name)
}


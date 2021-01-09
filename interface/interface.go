package main

import (
	"fmt"
)

type Sayer interface {
	say()
}

type Dog struct {

}

func (d Dog) say()  {
	fmt.Println("汪汪汪")
}

type Cat struct {

}

func (c Cat) say()  {
	fmt.Println("喵喵喵")
}

type Bird struct {

}

func (b Bird) say()  {
	fmt.Println("吱吱吱")
}

func do(arg Sayer)  {
	arg.say()
}

func main()  {
	c1 := Cat{}
	do(c1)
	d1 := Dog{}
	do(d1)
	b1 := Bird{}
	do(b1)

	var s Sayer
	s = c1
	s.say()

}


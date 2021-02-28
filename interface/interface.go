package main

import (
	"fmt"
)

type Sayer interface {
	say()
}

type Dog struct {
}

type Cat struct {
}

type Bird struct {
}

func (d *Dog) say() {
	fmt.Println("汪汪汪")
}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}

func (b Bird) say() {
	fmt.Println("吱吱吱")
}

func do(arg Sayer) {
	arg.say()
}

func main() {
	c1 := Cat{}
	do(c1)
	d1 := &Dog{}
	do(d1)
	b1 := Bird{}
	do(b1)

	var s Sayer
	s = c1
	s.say()

}

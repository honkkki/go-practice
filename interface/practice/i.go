package main

import (
	"fmt"
)

type inter interface {
	Read()
}

type anoInter interface {
	Read()
	Write()
}

type myint int

type mystring string

func (m myint) Read() {
	fmt.Println("int read")
}

func (m mystring) Read() {
	fmt.Println("string read")
}

func (m mystring) Write() {
	fmt.Println("string write")
}

func needInter(i inter) {
	fmt.Println(i)
}

func main() {
	var i1 inter
	var i2 anoInter

	var m1 myint
	var m2 mystring

	i1 = m1
	i2 = m2
	fmt.Println(i1)
	fmt.Println(i1 == i2)

	var i3 anoInter
	var i4 inter
	needInter(i3)
	fmt.Println(i3 == i4)
}

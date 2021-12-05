package main

import (
	"fmt"
)

type myints int

func (i *myints) Add() {
	fmt.Println(i)
}

func main() {
	var a myints = 1
	var i interface{} = a
	//i.(myints).Add()		// error
	m := i.(myints)
	m.Add()

}

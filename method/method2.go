package main

import "fmt"

type myint int
type myslice []int

func (m *myint) Set() {
	*m++
}

func (m *myslice) Append() {
	*m = append(*m, 1)
}

func (m myslice) Set() {
	m[0] = 2
}

func main() {
	var i myint = 1
	i.Set()
	fmt.Println(i)

	var s myslice
	s.Append()
	fmt.Println(s)

	s.Set()
	fmt.Println(s)
}

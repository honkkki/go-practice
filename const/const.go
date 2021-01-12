package main

import "fmt"

const (
	a = iota
	b
	c = 6
	d
	w
	q = iota
)

const (
	e = 1 << iota
	f
	g
)



func main()  {
	fmt.Println(a, b, c, d, w, q)
	fmt.Println(e,f,g)
}

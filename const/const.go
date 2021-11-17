package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c = 6    // 6
	d        // 6
	w        // 6
	q = iota // 5
	z
	s = iota
)

const (
	e = 1 << iota // 1
	f             // 2
	g             // 4
	x			  // 8
	y = e | g | x	// 13 1+4+8
)

func main() {
	fmt.Println(a, b, c, d, w, q, z, s)
	fmt.Println(e, f, g, x, y)
}

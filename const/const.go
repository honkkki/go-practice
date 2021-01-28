package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c = 6    // 6
	d        // 6
	w        // 6
	q = iota // 5
)

const (
	e = 1 << iota // 1
	f             // 2
	g             // 4
)

func main() {
	fmt.Println(a, b, c, d, w, q)
	fmt.Println(e, f, g)
}

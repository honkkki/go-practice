package main

import (
	"fmt"
)

func main() {
	var m map[string]int

	fmt.Println(m)
	//m["num"] = 1		// panic
	//fmt.Println(m)

	m = make(map[string]int, 256)
	m["num"] = 1
	fmt.Println(m)

	fmt.Println("-----------------")
	delete(m, "name")

	res := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println(res)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

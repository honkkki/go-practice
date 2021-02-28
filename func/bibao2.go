package main

import "fmt"

func addBase(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func main() {
	f := addBase(10)
	fmt.Println(f(1), f(2))
}

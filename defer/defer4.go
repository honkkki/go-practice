package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer func() {
		a++
		a++
		fmt.Println(a)
		fmt.Println(b)
	}()

	fmt.Println(a)
}

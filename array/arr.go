package main

import "fmt"

func main() {
	var a [2][3]int16
	fmt.Println(a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &a[0])
	fmt.Printf("%p\n", &a[0][0])
	fmt.Printf("%p\n", &a[0][1])
	fmt.Printf("%p\n", &a[0][2])
}

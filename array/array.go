package main

import "fmt"

func main() {
	arr1 := [2]string{"jisoo", "rose"}
	arr2 := [...]string{"jisoo", "rose"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Printf("%T \n", arr1)
}

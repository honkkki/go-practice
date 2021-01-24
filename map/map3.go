package main

import "fmt"

func main()  {
	var m map[string]int

	fmt.Println(m)
	//m["num"] = 1		// panic
	//fmt.Println(m)

	m = make(map[string]int, 256)
	m["num"] = 1
	fmt.Println(m)



}
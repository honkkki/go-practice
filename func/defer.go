package main

import "fmt"


func main() {
	//fmt.Println("start...")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println("end")

	x := 1
	defer func() {
		x += 1
		fmt.Println(x)
	}()
	x = 5

}

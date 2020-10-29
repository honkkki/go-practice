package calc

import "fmt"

var name = "add"

func init()  {
	fmt.Println("init")
	fmt.Println(name)
}

func Add(x, y int) int {
	return x + y
}
package calc

import "fmt"

var name = "add"

func init()  {
	fmt.Println("init")
}

func Add(x, y int) int {
	return x + y
}
package main

import (
	"fmt"
)

func main() {
	c := 3.54
	fmt.Println(int(c)) // 舍去小数部分
	arr := [3]int{1, 2, 3}
	fmt.Println(arr[:2])

}

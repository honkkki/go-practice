package main

import (
	"errors"
	"fmt"
)

func arraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}

	return sum
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arraySum(arr[:]))

	err := errors.New("foo")
	fmt.Printf("%T %v \n", err, err)
}

package main

import "fmt"

func tt() (a, _ int) {
	return 1, 2
}

func main() {
	fmt.Println(tt())
	s := make([]int, 10)
	s = []int{1, 2, 3}
	fmt.Println(s)

}

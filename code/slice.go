package main

import "fmt"

type TT struct {
	name string
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	s1 := s[:3:6] // 后面数字 cap容量
	fmt.Println(s1, cap(s1))

	t := TT{}
	fmt.Println(t == TT{})

}

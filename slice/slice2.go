package main

import (
	"fmt"
)

func main() {
	sli := make([]string, 0, 10)
	sli = append(sli, "hello")
	sli = append(sli, "hello")
	sli = append(sli, "hello")
	fmt.Println(sli, len(sli), cap(sli))

	var s1 []int
	fmt.Println(s1 == nil)
	s1 = append(s1, 1)
	fmt.Println(s1)

	s2 := []int{1, 2, 3}
	for k, v := range s2 {
		s2 = append(s2, 4)
		s2 = append(s2, 5)
		fmt.Println(k, v)
	}

}

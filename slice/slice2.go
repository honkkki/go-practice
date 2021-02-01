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

}

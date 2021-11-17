package main

import (
	"fmt"
)

type mm map[string]string

func main() {
	m := make(mm)
	m["name"] = "jisoo"
	fmt.Println(m)

	var num int
	num = -1
	un := uint(num)
	fmt.Printf("%T %v", un, un)

	go func() {
		go func() {
			return
		}()
	}()


}

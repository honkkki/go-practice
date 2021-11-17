package main

import "fmt"

func main() {
	m := make(map[int]int, 10)
	for i := 1; i<= 10; i++ {
		m[i] = i
	}

	fmt.Println(m)

	for k, v := range m {
		m[11] = 11
		fmt.Println(k, v)
	}

}

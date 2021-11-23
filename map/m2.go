package main

import "fmt"

func main() {
	m := make(map[int]int, 10)
	for i := 1; i <= 10; i++ {
		m[i] = i
	}

	fmt.Println(m)

	for k, v := range m {
		delete(m, 1)
		m[11] = 11
		fmt.Println(k, v)
	}

	fmt.Println("----------------------")
	s := []int{1, 2, 3}
	for k, v := range s { // 切片遍历时保持开始的状态
		s[0] = 2
		s = append(s, 4)
		fmt.Println(k, v)
	}

	fmt.Println(s)

}

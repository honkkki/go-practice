package main

import (
	"fmt"
	"sort"
)

func main() {
	// 数组排序
	arr := [5]int{5, 3, 4, 2, 1}
	sort.Ints(arr[:]) // 参数接收切片类型
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)

	s1 := arr[:]
	s := make([]int, len(arr))
	copy(s, s1)
	fmt.Println(s)
	fmt.Println(s1)

	fmt.Println("-------------------------")
	s1[0] = 2
	fmt.Println(s)
	fmt.Println(s1)

	fmt.Println("-------------------------")
	fmt.Println(removeI(s, 2))

	fmt.Println("-------------------------")
	sort.Sort(sort.Reverse(sort.IntSlice(s1)))
	fmt.Println(s1)

	fmt.Println("-------------------------")
	b := make([]byte, 3)
	copy(b, "123")
	fmt.Println(string(b))
}

func removeI(s []int, i int) []int {
	s = append(s[:i], s[i+1:]...)
	return s
}

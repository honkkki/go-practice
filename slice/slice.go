package main

import (
	"fmt"
	"sort"
)

func main() {
	// 切片与数组
	arr := [...]int{1, 2, 3, 4, 5}
	slice1 := arr[:3]
	slice2 := slice1[1:3]
	fmt.Println(cap(slice1), cap(slice2)) // 5 4

	// 切片扩容
	var s []int                      // 未初始化
	var s2 = []int{}                 // 已初始化
	fmt.Println(s == nil, s2 == nil) // true false 容量都为0 需要扩容
	s3 := []int{1, 2, 3}
	s = append(s, s3...)
	fmt.Println(s) // [1 2 3]

	// 切片copy
	s4 := make([]int, 3)
	fmt.Println(s4)
	copy(s4, s3) // 两个切片指向不同内存地址
	fmt.Printf("%p %p \n", s3, s4)
	s4[0] = 2
	fmt.Println(s4)

	// 切片删除元素
	s5 := []int{1, 2, 3, 4, 5}
	s5 = append(s5[:1], s5[2:]...)
	fmt.Println(s5)

	a := 1
	fmt.Printf("%T \n", fmt.Sprintf("%v", a)) // 转换为string类型

	arr2 := [...]int{1, 3, 2, 9, 6}
	sort.Ints(arr2[:])
	fmt.Println(arr2)
	fmt.Printf("%T \n", arr2)

}

package main

import (
	"fmt"
	"sort"
)

func main()  {
	// 数组排序
	arr := [5]int{5,3,4,2,1}
	sort.Ints(arr[:])		// 参数接收切片类型
	fmt.Println(arr)
	fmt.Printf("%T", arr)

	


}

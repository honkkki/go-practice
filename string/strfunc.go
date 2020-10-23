package main

import (
	"fmt"
	"strings"
)

// 字符串方法
func main()  {
	str := "jisoo"
	fmt.Println(len(str))
	fmt.Printf("%s-good\n", str)

	str1 := "hello goland"
	arr := strings.Split(str1, " ")
	fmt.Println(arr)
	fmt.Printf("%T", arr)


	
}
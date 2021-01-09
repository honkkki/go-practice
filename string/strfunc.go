package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串方法
func main()  {
	str := "jisoo"
	fmt.Println(len(str))
	fmt.Printf("%s-good \n", str)

	// 字符串转为切片
	str1 := "hello goland"
	arr := strings.Split(str1, " ")
	fmt.Println(arr)
	fmt.Printf("%T \n", arr)

	// 字符串的索引
	s := "supreme"
	sli := make([]string, 0, 10)
	index := strings.Index(s, "p")
	fmt.Println(index)
	fmt.Println(string(s[2]))
	res := append(sli, s[2:])
	fmt.Println(res)

	// 一个中文在utf8下占三个字节 利用索引取值需要考虑字节长度
	s1 := "哈哈啦"
	fmt.Println(strings.Index(s1, "啦"))		// 6

	// 整型转为字符串
	i := 1
	intToStr := strconv.Itoa(i)
	fmt.Println(intToStr)
}



package main

import (
	"crypto/sha256"
	"fmt"
)

// 字符操作
func main() {
	var c1 byte = 'c'
	var c2 rune = 'c'
	var c3 rune = '中'
	fmt.Println(c1, c2, c3)
	fmt.Printf("%T %T \n", c1, c2)
	fmt.Println("---------------------------")

	s1 := "jisoo"
	// 字符串是只读的 无法通过下标修改字符进行修改 -> 转成[]byte再修改
	fmt.Println(s1[0]) // 106 byte类型

	s1Arr := []byte(s1)
	fmt.Println(s1Arr)
	s1Arr[0] = 'J'
	fmt.Println(string(s1Arr[0]))
	fmt.Println(string(s1Arr))
	fmt.Println(len(string(s1Arr))) // 字节数就是长度

	str := "x"
	str1 := "X"
	fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	fmt.Printf("%x\n", sha256.Sum256([]byte(str1)))

}

package main

import "fmt"

// 字符操作
func main() {
	var c1 byte = 'c'
	var c2 rune = 'c'
	var c3 rune = '中'
	fmt.Println(c1, c2, c3)
	fmt.Printf("%T %T \n", c1, c2)

	s1 := "jisoo"
	s1Arr := []byte(s1)
	fmt.Println(s1Arr)
	fmt.Println(string(s1Arr[0]))

}

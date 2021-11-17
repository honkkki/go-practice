package main

import "fmt"

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	s = remove(s, 1)
	fmt.Println(s)
	s = reverse(s)
	fmt.Println(s)

	// string replace
	str := "abc"
	strBytes := []byte(str)
	strBytes[0] = 'b'
	str = string(strBytes)
	fmt.Println(str)

	arr := [1]int{1}
	arr2 := [1]int{1}
	fmt.Println(arr==arr2)		// 数组可以比较

	//s1 := []int{1}
	//s2 := []int{1}
	//fmt.Println(s1==s2)		// error: 切片只能跟nil比较 map也一样

}

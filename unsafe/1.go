package main

import (
	"fmt"
	"unsafe"
)

func main() {
	m := make(map[string]string)
	fmt.Println(unsafe.Sizeof(m)) // map初始化返回的是指针

	s := []int{}
	fmt.Println(unsafe.Sizeof(s)) // slice header struct

	str := "golang"
	fmt.Println(unsafe.Sizeof(str)) // string header struct

}

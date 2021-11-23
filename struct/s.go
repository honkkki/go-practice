package main

import (
	"fmt"
	"unsafe"
)

type Flag struct {
	num1 int16 // 2
	num2 int32 // 4
}

type demo1 struct {
	a int16  // 2
	b int32  // 4
	s string // 16
}

type demo2 struct {
	a int16  // 2
	s string // 16
	b int32  // 4
}

type demo3 struct {
	a int8
	c int32
	b int16
	d int64
}

func main() {
	fmt.Println(unsafe.Sizeof(Flag{})) // 8 内存对齐
	fmt.Println(unsafe.Sizeof(demo1{}))
	fmt.Println("align:", unsafe.Alignof(demo1{}))
	fmt.Println(unsafe.Sizeof(demo2{}))
	fmt.Println(unsafe.Sizeof(demo3{}))
	fmt.Println("align:", unsafe.Alignof(demo3{}))
}

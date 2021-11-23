package main

import (
	"fmt"
	"unsafe"
)

type Pet struct {
	Name string
	Age  int
}

func main() {
	a := 1
	switch {
	case a == 0:
		fmt.Println("a=0")
	case a == 1:
		fmt.Println("a=1")
	}

	switch aa := 100; {
	case aa == 1:
		fmt.Println("aa=1")
	case aa == 100:
		fmt.Println("aa=100")
	}

	fmt.Println(unsafe.Sizeof(a))

	var num int32 = 1
	str := "hello"
	fmt.Println(unsafe.Sizeof(num))
	fmt.Println(unsafe.Sizeof(str))

	pet := Pet{}
	fmt.Println(unsafe.Sizeof(pet))  // 24 string 16+int64 8
	fmt.Println(unsafe.Sizeof(&pet)) // 8

}

package main

import (
	"fmt"
	"unsafe"
)

// 嵌套结构体
type Address struct {
	Province string
	City     string
}

type Human struct {
	Name string
	Age  int
	*Address
}

func main() {
	human := Human{
		Name: "jisoo",
		Age:  20,
		Address: &Address{
			"gd",
			"sz",
		},
	}

	fmt.Printf("%#v \n", human)
	fmt.Printf("address : %#v \n", human.Address)
	fmt.Println(human.Province)

	h1 := Human{}
	fmt.Println(h1==Human{})		// 空结构体
	fmt.Println(unsafe.Sizeof(h1))
	fmt.Println(unsafe.Sizeof(human))

}

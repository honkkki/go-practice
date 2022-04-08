package main

import (
	"fmt"
	"unsafe"
)

func main() {
	m := make(map[string]string)
	m["name"] = "tom"
	fmt.Println(unsafe.Sizeof(m))
	m["age"] = "10"
	fmt.Println(unsafe.Sizeof(m))
}

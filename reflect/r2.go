package main

import (
	"fmt"
	"reflect"
)

func main() {
	// change slice by reflect
	s := []int{1, 2, 3}
	v := reflect.ValueOf(s)
	v.Index(0).SetInt(2)
	fmt.Println(s)

	m := map[string]string{
		"name": "tom",
	}
	vm := reflect.ValueOf(m)
	vm.SetMapIndex(reflect.ValueOf("name"), reflect.ValueOf("karina"))
	fmt.Println(m)
}

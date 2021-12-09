package main

import (
	"fmt"
	"reflect"
)

// 类型断言
func main() {
	a := 1
	t := reflect.TypeOf(a)
	fmt.Println(t.Name())
	fmt.Printf("%T\n", t)
	fmt.Println("----------------------------------")

	// 空接口判断值类型 接口类型转为具体值类型
	var itf interface{}
	var num int64 = 100
	itf = num
	fmt.Printf("%T\n", itf)
	v, ok := itf.(int)
	if ok {
		fmt.Println(v)
		fmt.Printf("%T\n", v) // int
	} else {
		fmt.Println(v)
		fmt.Printf("%T\n", v) // int
	}

	v1 := itf.(int64)
	fmt.Println(v1)
	fmt.Println(reflect.TypeOf(v1))

	fmt.Println("----------------------------------")
	// switch判断接口类型
	switch val := itf.(type) {
	case string:
		fmt.Println("is string", val)
		break
	case int:
		fmt.Println("is int", val)
		break
	}

}

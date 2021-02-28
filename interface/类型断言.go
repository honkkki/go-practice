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

	// 空接口判断值类型 接口类型转为具体值类型
	var itf interface{}
	itf = 1
	v := itf.(int)
	fmt.Println(v)
	fmt.Printf("%T\n", v)

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

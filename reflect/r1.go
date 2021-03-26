package main

import (
	"fmt"
	"reflect"
	"strings"
)

func ref(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println(t.Name())

	fmt.Println(t.Kind())

	if t.Kind() == reflect.Int8 {
		fmt.Printf("is: %v\n", t.Kind().String())
	}

	v := reflect.ValueOf(a)
	fmt.Println(v)
	fmt.Println(v.Kind())
	b := strings.Contains(v.Kind().String(), "int")
	fmt.Println(b)
	fmt.Println(v.Int())

}

func main() {
	var a int8 = 1
	ref(a)

	// reflect set
	v := reflect.ValueOf(&a)	// 必须传地址才能改变本身的值
	v.Elem().SetInt(10)		// 获取指针地址保存的值再设置新的值
	fmt.Println(a)

}

package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
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
	fmt.Printf("%T\n", v)
	fmt.Println(v.Int())
	fmt.Println("--------------------------")
}

func main() {
	var a int8 = 1
	ref(a)

	// reflect set
	v := reflect.ValueOf(&a) // 必须传地址才能改变本身的值
	//v.Elem().Set(reflect.ValueOf(int8(10)))
	v.Elem().SetInt(10) // 获取指针地址保存的值再设置新的值
	fmt.Println(a)

	var w io.Writer
	fmt.Println(reflect.TypeOf(w))
	w = os.Stdout
	fmt.Println(reflect.TypeOf(w))
}

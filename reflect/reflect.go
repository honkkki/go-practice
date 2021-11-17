package main

import (
	"fmt"
	"reflect"
)

// 反射
func main()  {
	var a int8 = 1
	// 获取类型
	res := reflect.TypeOf(a)
	fmt.Println(res)
	fmt.Printf("%T \n", res)
	// 获取值
	data := reflect.ValueOf(a)
	// 获取变量的种类
	fmt.Println(data.Kind())
	fmt.Println(data)
	fmt.Println(data.Int())
	d := data.Int() + 1
	fmt.Println(d)



}

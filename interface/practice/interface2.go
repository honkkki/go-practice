package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 空接口
	var x interface{}
	x = 1
	fmt.Printf("%T\n", x)

	// 类型断言
	res, ok := x.(int)
	if ok {
		fmt.Println(res)
	} else {
		fmt.Println("error type")
	}

	// 空接口的使用
	var slice = make([]interface{}, 5)
	slice[0] = 1
	slice[1] = "jisoo"
	fmt.Println(slice)

	fmt.Println(runtime.NumCPU())

}

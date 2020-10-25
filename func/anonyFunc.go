package main

import (
	"fmt"
)

// 匿名函数

func testFunc1(name string) {
	// 闭包 匿名函数使用了外层的变量
	func() {
		fmt.Println(name)
	}()
}

func main() {
	func1 := func() {
		fmt.Println("anonymous")
	}

	func1()

	// 立即执行
	func() {
		fmt.Println("anonymous")
	}()

	testFunc1("golang")
}

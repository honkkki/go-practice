package main

import "fmt"

func a() {
	defer func() {
		err := recover() // 搭配defer使用
		if err != nil {
			fmt.Println("func error")
		}
	}()

	panic("error")
}

func b() {
	fmt.Println("b func")
}

func main() {
	a()
	b()
}

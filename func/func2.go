package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	f := add
	fmt.Printf("%T \n", f)

	i := 1
	defer func() {
		fmt.Println(i)
	}() // defer后匿名函数 最后执行的时候 i=100
	i = 100
	fmt.Println(i)

	// 区别于
	// i := 1
	//defer func(i int) {
	//	fmt.Println(i)
	//}(i)							// 这时i已经传入 i=1
	//i = 100
	//fmt.Println(i)

	fmt.Println("---------------------------------")

	for i := 0; i < 5; i++ {
		//i := i
		func() {
			println(i)
		}()
	}

}

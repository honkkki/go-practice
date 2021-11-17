package main

import "fmt"

func main() {
	arr1 := [2]string{"jisoo", "rose"}
	arr2 := [...]string{"jisoo", "rose"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Printf("%T \n", arr1)

	arr3 := [3]interface{}{
		1, "a", true,
	}

	fmt.Println(arr3)
	fmt.Printf("%T \n", arr3)

	a := [3]int{}
	fmt.Println(a)

	var arr [3]int
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println(arr)


}

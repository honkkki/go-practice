package main

import "fmt"

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}

	return arr
}

func main() {
	arr := []int{5, 6, 3, 1, 9, 7}
	fmt.Println(arr)
	insertSort(arr)
	fmt.Println(arr)
}

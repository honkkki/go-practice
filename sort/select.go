package main

import "fmt"

// 从数中选出一个最小的放到前面 （选一个数跟后面的数比较）
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := []int{5, 6, 3, 1, 9, 7}
	fmt.Println(arr)
	selectSort(arr)
	fmt.Println(arr)
}

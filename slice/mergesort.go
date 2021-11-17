package main

import "fmt"

/*
	两个有序数组合并成一个有序数组 Golang实现
*/
func main() {
	var a = []int{1, 3, 5, 7, 9}
	var b = []int{0, 2, 4, 6, 8}

	c := sortArr(a, b)

	for i, v := range c {
		fmt.Println(i, ":", v)
	}
}

func sortArr(a, b []int) []int {

	//计算数组长度
	alen := len(a)
	blen := len(b)
	clen := alen + blen

	//创建结果集数组
	c := make([]int, clen)

	//默认数组从0开始
	aIndex := 0
	bIndex := 0
	cIndex := 0

	//开始遍历数组
	for aIndex < alen && bIndex < blen {
		if a[aIndex] < b[bIndex] {
			c[cIndex] = a[aIndex]
			cIndex++
			aIndex++
		} else {
			c[cIndex] = b[bIndex]
			cIndex++
			bIndex++
		}
	}

	// 两数组长度不同 当其中一个数组遍历完,另一个数组可能还没遍历完 将另一个数组遍历完
	for aIndex < alen {
		c[cIndex] = a[aIndex]
		cIndex++
		aIndex++
	}
	for bIndex < blen {
		c[cIndex] = b[bIndex]
		cIndex++
		bIndex++
	}

	return c
}

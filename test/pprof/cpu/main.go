package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func quickSort(nums []int) []int {
	var left []int
	var right []int
	var res []int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[0] {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	res = append(left, nums[0])
	res = append(res, right...)
	return res
}

func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0777)
	defer f.Close()
	err := pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		bubbleSort(nums)
		quickSort(nums2)
		n *= 10
	}
}

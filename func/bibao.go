package main

import (
	"fmt"
	"strings"
)

func makeSuffix(name string, suffix string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}

	return name
}

// 使用闭包
func makeSuffixClosure(name string) func(suffix string) string {
	return func(suffix string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func main() {
	res := makeSuffix("test", ".txt")
	fmt.Println(res)

	closureFunc := makeSuffixClosure("test")
	ret := closureFunc(".txt")
	fmt.Println(ret)
}

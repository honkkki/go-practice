package main

import (
	"fmt"
	"strings"
)

func main() {
	// 统计一个字符串单词出现次数
	m1 := make(map[string]int)
	str := "hello golang"
	words := strings.Split(str, " ")

	for _, word := range words {
		for _, v := range word {
			//fmt.Println(string(v))
			if times, ok := m1[string(v)]; ok {
				m1[string(v)] = times + 1
			} else {
				m1[string(v)] = 1
			}
		}
	}

	for k, v := range m1 {
		fmt.Printf("%s => %v \n", k, v)
	}
}

package main

import (
	"fmt"
	"strings"
)

func wordCount(str string) map[string]int {
	res := make(map[string]int, 100)
	words := strings.Split(str, " ")
	for _, v := range words {
		count, ok := res[v]
		if !ok {
			res[v] = 1
		} else {
			res[v] = count + 1
		}
	}

	return res
}

func main() {
	res := wordCount("hello golang golang good")
	fmt.Println(res)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	strs := []string{"one", "two", "three"}

	// 闭包引用的外部for循环的s
	for _, s := range strs {
		go func() {
			fmt.Printf("%s ", s)
		}()
	}

	time.Sleep(1 * time.Second)
}

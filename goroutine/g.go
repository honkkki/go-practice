package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()

	runtime.Gosched()
	runtime.GC()
	fmt.Println("done")
}

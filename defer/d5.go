package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		i := i
		defer func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(3 * time.Second)
}

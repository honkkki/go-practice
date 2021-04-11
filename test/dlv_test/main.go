package main

import (
	"fmt"
	"time"
)

func main() {
	for  {
		var i int
		var currentTime time.Time
		i++
		currentTime = time.Now()
		fmt.Printf("[%d]-----, %v\n ", i, currentTime)
		time.Sleep(time.Second * 5)
	}

}

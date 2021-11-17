package main

import (
	"fmt"
	"time"
)

// go实现定时器的功能
func main()  {
	ticker := time.Tick(time.Second)
	for v := range ticker {
		fmt.Println("每秒hello一次", v)
	}

	

}

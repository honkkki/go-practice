package main

import (
	"log"
	"time"
)

func main() {
	// goroutine自己捕获自己的panic
	defer func() {
		if err := recover(); err != nil {
			log.Println("catch panic:", err)
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch panic:", err)
			}
		}()

		panic("goroutine error")
	}()

	//panic("error")
	time.Sleep(time.Second * 5)
}

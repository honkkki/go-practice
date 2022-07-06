package main

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

var a int
var mu sync.Mutex

func tick(name string) func() {
	return func() {

	}
}

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println("tick 1s")
		mu.Lock()
		defer mu.Unlock()
		a++
		fmt.Println(a)
	})

	c.AddFunc("* * * * *", func() {
		fmt.Println("1 minute")
		mu.Lock()
		defer mu.Unlock()
		a += 10
		fmt.Println(a)
	})

	c.Start()

	select {}
}

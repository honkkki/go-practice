package main

import (
	"fmt"
	"sync"
)

var m = make(map[string]int)
var wg sync.WaitGroup
var mu sync.Mutex
var so sync.Once

func InitMap()  {
	so.Do(func() {
		m["num"] = 0
	})
}

func con()  {
	InitMap()		// 只执行一次
	defer func() {
		wg.Done()
		mu.Unlock()
	}()
	mu.Lock()
	m["num"]++
}


func main()  {
	for i:=0;i<1000;i++ {
		wg.Add(1)
		go con()
	}

	wg.Wait()
	fmt.Println(m["num"])
}

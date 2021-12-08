package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func worker() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		runtime.Gosched()
		wg.Done()
	}
}

func main() {
	wg.Add(10)
	go worker()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		runtime.Gosched()
	}

	runtime.GC()
	wg.Wait()

	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("GOOS:", runtime.GOOS)

}

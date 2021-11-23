package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/ants/v2"
)

func worker() func() {
	return func() {

	}
}

func main() {
	p, _ := ants.NewPool(10000)

	defer p.Release()

	var wg sync.WaitGroup

	wg.Add(1)
	err := p.Submit(func() {
		fmt.Println("worker")
		wg.Done()
	})

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	pf, _ := ants.NewPoolWithFunc(10000, func(i interface{}) {
		n := i.(int)
		fmt.Println(n)
		wg.Done()
	})
	defer pf.Release()

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		pf.Invoke(i)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs chan int, results chan int)  {
	for j := range jobs{
		fmt.Printf("%d worker start job %d \n", id, j)
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("%d worker finish job %d \n", id, j)
		results <- j
		//wg.Done()
	}

}

func main()  {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	//wg.Add(5)

	for w:=1;w<=3;w++ {
		go worker(w, jobs, results)
	}

	for j:=1;j<=5;j++ {
		jobs <- j
	}

	close(jobs)

	// 为了等goroutine执行完才退出主线程
	for i:=0;i<5;i++ {
		<- results
	}

	//wg.Wait()

}

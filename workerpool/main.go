package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	Id     int
	Number int
}

type Result struct {
	*Job
	Sum int
}

func calc(job *Job, result chan *Result) {
	var sum int
	number := job.Number
	for number != 0 {
		tmp := number % 10
		sum += tmp
		number /= 10
	}

	r := &Result{
		job,
		sum,
	}

	result <- r
}

func worker(jobChan chan *Job, resChan chan *Result) {
	for job := range jobChan {
		calc(job, resChan)
	}

}

// 创建num个goroutine
func workerPool(num int, jobChan chan *Job, resChan chan *Result) {
	for i := 0; i < num; i++ {
		go worker(jobChan, resChan)
	}

}

func printResult(resChan chan *Result) {
	for res := range resChan {
		fmt.Printf("job id: %v, number: %v, sum: %v\n", res.Job.Id, res.Job.Number, res.Sum)
	}

}

func main() {
	var id int
	rand.Seed(time.Now().Unix())
	jobChan := make(chan *Job, 1000)
	resChan := make(chan *Result, 1000)
	workerPool(200, jobChan, resChan)
	go printResult(resChan)

	for {
		number := rand.Int31()
		id++
		job := &Job{
			Id:     id,
			Number: int(number),
		}

		jobChan <- job

		time.Sleep(1 * time.Second)
	}

}

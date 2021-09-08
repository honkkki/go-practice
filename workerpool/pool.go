package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Task struct {
	f    func()
}

func NewTask(f func()) *Task {
	return &Task{
		f: f,
	}
}

func (t *Task) Exec() {
	t.f()
}

type Pool struct {
	EntryChan chan *Task
	workerNum int
	JobsChan  chan *Task
}

func NewPool(workerNum int) *Pool {
	return &Pool{
		EntryChan: make(chan *Task, 100),
		workerNum: workerNum,
		JobsChan:  make(chan *Task, 100),
	}
}

func (p *Pool) worker(workerID int) {
	for task := range p.JobsChan {
		task.Exec()
		fmt.Println("worker id:", workerID, "finish work...")
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}

	for task := range p.EntryChan {
		p.JobsChan <- task
	}
}

func main() {
	pool := NewPool(10)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		for {
			select {
			case <-sig:
				// 释放资源
				close(pool.JobsChan)
				close(pool.EntryChan)
				fmt.Println("progress exit!")
			default:
				task := NewTask(func() {
					fmt.Println("task:", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
				})
				pool.EntryChan <- task
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	pool.Run()
}

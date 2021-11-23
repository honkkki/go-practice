package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func play(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}

}

func main() {
	table := make(chan *Ball)
	go play("pong", table)
	go play("ping", table)
	table <- new(Ball)
	time.Sleep(3 * time.Second)
	<-table
	fmt.Println("finish")
}

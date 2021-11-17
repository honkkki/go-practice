package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var (
	randNum int
)

func worker(ctx context.Context) {
	ch := make(chan struct{}, 1)
	fmt.Println(randNum)
	if randNum == 1 {
		ch <- struct{}{}
	}

	select {
	case <-ctx.Done():
		fmt.Println("overtime")
	case <-ch:
		fmt.Println("success")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	randNum = rand.Intn(3)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	worker(ctx)
	fmt.Println("over")
}

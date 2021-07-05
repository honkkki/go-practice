package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	g := new(errgroup.Group)
	// run a goroutine
	// 执行一批goroutine获得第一个error. 例如http.Get()
	for i:=0;i<5;i++ {
		g.Go(func() error {
			return errors.New("error in goroutine")
		})
	}

	if err := g.Wait(); err != nil {
		log.Println(err)
	} else {
		fmt.Println("success")
	}
}

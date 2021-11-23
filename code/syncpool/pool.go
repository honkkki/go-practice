package main

import (
	"fmt"
	"sync"
	"time"
)

type Logger struct {
	time    string
	content string
}

var sp = sync.Pool{
	New: func() interface{} {
		return new(Logger)
	},
}

func main() {
	cached := sp.Get().(*Logger)
	cached.time = time.Now().Format("2006-01-02 15:04:05")
	cached.content = "first content"
	fmt.Printf("%#v\n", cached)

	sp.Put(cached)
	cached2 := sp.Get().(*Logger)
	fmt.Printf("%#v\n", cached2)

	cached.time = ""
	cached.content = ""
	sp.Put(cached)
	cached3 := sp.Get().(*Logger)
	fmt.Printf("%#v\n", cached3)

}

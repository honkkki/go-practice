package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type entry struct {
	res   result
	ready chan struct{}
}

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
}

type result struct {
	value interface{}
	err   error
}

func NewMemo(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]

	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	}

}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := NewMemo(httpGetBody)
	incomingURLs := []string{
		"https://golang.org",
		"https://www.baidu.com",
		"https://gobyexample.com",
		"https://books.studygolang.com",
		"https://golang.org",
		"https://gobyexample.com",
	}

	var wg sync.WaitGroup

	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Printf("get data fail: %v\n", err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	num    int
	mlock  sync.Mutex   // 互斥锁
	rwlock sync.RWMutex // 读写锁
	wg1    sync.WaitGroup
	m      sync.Map
)

// 使用读写锁 应用在读多写少的场景
func read() {
	//mlock.Lock()
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	//mlock.Unlock()
	rwlock.RUnlock()

	wg1.Done()
}

func write(i int) {
	//mlock.Lock()
	rwlock.Lock()
	num++
	m.Store(i, time.Now().Nanosecond())
	time.Sleep(time.Millisecond * 10)
	//mlock.Unlock()
	rwlock.Unlock()

	wg1.Done()
}

func main() {
	startTime := time.Now()

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go write(i)
	}

	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go read()
	}

	wg1.Wait()
	fmt.Println(time.Now().Sub(startTime))
	value, _ := m.Load(1)
	value2, _ := m.Load(2)
	fmt.Println(value)
	fmt.Println(value2)
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)

// dirFiles get dir/files[].
// 获取文件夹下所有文件
func dirFiles(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read dir error: %v\n", err)
		return nil
	}

	return files
}

// walkDir 遍历目录计算文件总大小
func walkDir(dir string, fileSize chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	files := dirFiles(dir)
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, file.Name())
			go walkDir(subDir, fileSize, wg)
		} else {
			fileSize <- file.Size()
		}
	}
}

func printSize(nFiles, nBytes int64) {
	fmt.Printf("files: %d, size: %.1fmb\n", nFiles, float64(nBytes/1e6))
}

func main() {
	// 输入-v表示打印计算详情
	showDetail := flag.Bool("v", false, "show detail message")
	flag.Parse()
	// tick nil now. if show detail it will be a tick
	var tick <-chan time.Time
	if *showDetail {
		tick = time.Tick(100 * time.Millisecond)
	}

	paths := flag.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}

	fileSize := make(chan int64, 100)
	var wg sync.WaitGroup

	go func() {
		for _, path := range paths {
			wg.Add(1)
			go walkDir(path, fileSize, &wg)
		}

		wg.Wait()
		close(fileSize)
	}()

	var nBytes, nFiles int64
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size
		case <-tick:
			printSize(nFiles, nBytes)
		}

	}

	//for size := range fileSize {
	//	nFiles++
	//	nBytes+=size
	//	printSize(nFiles, nBytes)
	//}

	printSize(nFiles, nBytes)
}

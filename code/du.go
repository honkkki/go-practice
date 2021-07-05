package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var sema = make(chan struct{}, 20)

// get dir/files[].
// 获取文件夹下所有文件
func dirFiles(dir string) []os.FileInfo {
	sema<- struct{}{}
	defer func() {
		<-sema
	}()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read dir error: %v", err)
		return nil
	}

	return files
}

// 遍历目录计算文件总大小
func walkDir(dir string, fileSize chan<- int64, wg *sync.WaitGroup)  {
	defer wg.Done()
	files := dirFiles(dir)
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, file.Name())
			go walkDir(subDir, fileSize, wg)
		} else {
			fileSize<-file.Size()
		}
	}
}

func printSize(nfiles, nbytes int64)  {
	fmt.Printf("files: %d, size: %.1fmb\n", nfiles, float64(nbytes/1e6))
}

func main()  {
	flag.Parse()
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

	var nbytes, nfiles int64
	for size := range fileSize {
		nfiles++
		nbytes+=size
	}

	printSize(nfiles, nbytes)

	//var tick <-chan time.Time
	//showDetail := flag.Bool("v", false, "show detail message")
	//if *showDetail {
	//	tick = time.Tick(500*time.Millisecond)
	//}
	//
	//select {
	//case <-tick:
	//default:
	//
	//}

}

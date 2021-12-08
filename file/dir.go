package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func RangeDir(dir string, level int) {
	files, _ := ioutil.ReadDir(dir)

	s := "|---"
	for i := 0; i < level; i++ {
		s += "|---"
	}

	for _, file := range files {
		fmt.Printf("%s%s\n", s, file.Name())
		if file.IsDir() {
			filePath := filepath.Join(dir, file.Name())
			RangeDir(filePath, level+1)
		}
	}
}

func main() {
	baseDir := "/mnt/d/ins-img"
	RangeDir(baseDir, 0)
}

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("./1.go")
	fmt.Println(path)
	fmt.Println(filepath.Dir(path))
}

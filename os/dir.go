package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dir := "."
	file, _ := ioutil.ReadDir(dir)
	fmt.Println(file)
	for _, fi := range file {
		fmt.Println(fi.Name())
	}
}

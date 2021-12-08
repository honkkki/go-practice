package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("1.go")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024, 1024)

	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
	}

}

package main

import (
	"os"
)

func main() {
	var buf [8]byte
	os.Stdin.Read(buf[:])
	//fmt.Printf("%s", string(buf[:]))
	os.Stdout.Write(buf[:])

}

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func newProcess() {
	cmd := exec.Command("./cli")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("son:", cmd.Process.Pid)

	outPut, err := exec.Command("./cli").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(outPut))
}

func main() {
	fmt.Println("main:", os.Getpid())
	newProcess()
	for {

	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func catFile(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	// 获取命令行参数
	flag.Parse()
	if flag.NArg() == 0 {
		catFile(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		catFile(bufio.NewReader(f))
	}

}

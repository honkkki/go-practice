package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

func main()  {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		//b, _ := ioutil.ReadAll(resp.Body)

		// 写入文件
		file, _ := os.Create("text.txt")
		buf := bufio.NewWriter(file)
		io.Copy(buf, resp.Body)
		//io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		buf.Flush()
		//fmt.Println(string(b))
	}

}

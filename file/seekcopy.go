package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// 断点续传
func main() {
	srcFile := "./m.jpg"
	dstFile := "./m1.jpg"
	tmpFile := "./tmp.txt"
	f1, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f3, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		f1.Close()
		f2.Close()
		f3.Close()
	}()

	var haveWrite int
	b, _ := ioutil.ReadAll(f3)
	if len(b) > 0 {
		haveWrite, _ = strconv.Atoi(string(b))
	}

	// read from f1 write to f2
	for {
		buf := make([]byte, 1024, 1024)
		f1.Seek(int64(haveWrite), io.SeekStart)
		n, err := f1.Read(buf)
		if err != nil {
			if err == io.EOF {
				f3.Truncate(0)
				fmt.Println("write finish bytes:", haveWrite)
				break
			} else {
				log.Fatal(err)
			}
		}

		f2.Seek(int64(haveWrite), io.SeekStart)
		_, err = f2.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
		haveWrite += n
		f3.Seek(0, io.SeekStart)
		_, err = f3.WriteString(strconv.Itoa(haveWrite))
		if err != nil {
			log.Fatal(err)
		}

		//if haveWrite > 10000 {
		//	log.Fatal("shut down for fake error.")
		//}
	}

}

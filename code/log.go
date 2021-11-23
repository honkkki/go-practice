package main

import "fmt"

func main() {
	// log打印文件与行号
	//log.SetFlags(log.Lshortfile|log.LstdFlags)
	//log.Println("hello")

	ch := make(chan struct{})

	go func() {
		for c := range ch { // error 阻塞
			fmt.Println(c)
		}
	}()

	ch <- struct{}{}

	//close(ch)

	select {} // error no case阻塞
	//go func() {
	//	ch <- struct{}{}
	//	close(ch)
	//}()
	//
	//for s := range ch {
	//	fmt.Println(s)
	//}
}

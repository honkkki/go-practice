package main

import "fmt"

// 往ch1中塞值
func f1(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}

	close(ch1)
}

// 从ch1取值并往ch2中塞值
func f2(ch1 <-chan int, ch2 chan<- int) {
	for num := range ch1 {
		ch2 <- num * num
	}

	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go f1(ch1)
	go f2(ch1, ch2)

	for num := range ch2 {
		fmt.Println(num)
	}

}

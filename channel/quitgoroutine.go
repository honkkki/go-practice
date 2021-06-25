package main

import "fmt"

func workerG()  {

}


func main()  {
	go workerG()

	fmt.Println("finish")
}

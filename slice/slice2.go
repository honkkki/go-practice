package main

import (
	"fmt"
)

func main()  {
	sli := make([]string, 0, 10)
	sli = append(sli, "hello")
	sli = append(sli, "hello")
	sli = append(sli, "hello")
	fmt.Println(sli, len(sli), cap(sli))


}

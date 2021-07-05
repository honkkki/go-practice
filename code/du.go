package main

import (
	"flag"
	"time"
)

func main()  {
	var tick <-chan time.Time
	showDetail := flag.Bool("v", false, "show detail message")
	if *showDetail {
		tick = time.Tick(500*time.Millisecond)
	}



}

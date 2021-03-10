package main

import (
	"go-practice/interface/logger"
	"time"
)

func main()  {
	logger.InitLogger()
	for {
		logger.Debug("user server is running")
		time.Sleep(time.Second)
	}
}

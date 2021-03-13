package main

import (
	"go-practice/logger/logger"
	"time"
)

func main()  {
	logger.InitLogger()
	for {
		logger.Debug("user server is running")
		time.Sleep(time.Second)
	}
}

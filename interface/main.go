package main

import (
	"go-practice/interface/logger"
)

func main()  {
	logger.InitLogger()
	logger.Debug("new debug")
}

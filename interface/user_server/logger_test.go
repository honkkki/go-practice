package user_server

import (
	"fmt"
	"go-practice/interface/logger"
	"testing"
)

func TestFileLogger(t *testing.T) {
	fmt.Println("test file logger")
	log := logger.NewFileLogger("../log_file", "test")
	defer log.Close()
	log.Debug("log debug")
	log.Info("log info")
	log.Error("log error")
	log.Info([]string{
		"hello", "golang",
	})
}

func TestConsoleLogger(t *testing.T) {
	log := logger.NewConsoleLogger()
	log.Debug("log console debug")
}
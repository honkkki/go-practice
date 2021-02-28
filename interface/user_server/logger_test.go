package user_server

import (
	"fmt"
	"go-practice/interface/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	fmt.Println("test function")
	l := logger.NewFileLogger("../log_file", "test")
	l.Info("hello golang")
}

package logger

import "strconv"

var logger LoggerInterface
var config = make(map[string]string, 100)

func init() {
	config["type"] = "file"
	config["fp"] = "./log_file"
	config["fn"] = "test"
	config["log_chan_size"] = "1000"
}

func InitLogger() {
	switch config["type"] {
	case "file":
		size, err := strconv.Atoi(config["log_chan_size"])
		if err != nil {
			size = 1000
		}
		logger = NewFileLogger(config["fp"], config["fn"], size)
	case "console":
		logger = NewConsoleLogger()
	}
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Close() {
	logger.Close()
}

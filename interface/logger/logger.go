package logger

var logger LoggerInterface
var config = make(map[string]string, 100)

func init()  {
	config["type"] = "file"
	config["fp"] = "./log_file"
	config["fn"] = "test"
}

func InitLogger()  {
	switch config["type"] {
	case "file":
		logger = NewFileLogger(config["fp"], config["fn"])
	case "console":
		logger = NewConsoleLogger()
	}
}

func Debug(args ...interface{})  {
	logger.Debug(args...)
}


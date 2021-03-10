package logger

import "os"

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger() LoggerInterface {
	consoleLogger := &ConsoleLogger{}
	return consoleLogger
}

func (f *ConsoleLogger) SetLevel(level int) {
	if level < DEBUG || level > ERROR {
		level = DEBUG
	}
	f.level = level
}

func (f *ConsoleLogger) Debug(args ...interface{}) {
	f.SetLevel(DEBUG)
	WriteLog(os.Stdout, f.level, args...)
}

func (f *ConsoleLogger) Info(args ...interface{}) {
	f.SetLevel(INFO)
	WriteLog(os.Stdout, f.level, args...)
}

func (f *ConsoleLogger) Error(args ...interface{}) {
	f.SetLevel(ERROR)
	WriteLog(os.Stdout, f.level, args...)
}

func (f *ConsoleLogger) Close() {
}

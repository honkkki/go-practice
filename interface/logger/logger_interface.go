package logger

type LoggerInterface interface {
	SetLevel(level int)
	LogPrefix()
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Close()
}

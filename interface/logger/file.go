package logger

import (
	"fmt"
	"os"
)

type FileLogger struct {
	level    int
	filePath string
	fileName string
	file     *os.File
}

func NewFileLogger(fp, fn string) LoggerInterface {
	fileLogger := &FileLogger{
		filePath: fp,
		fileName: fn,
	}

	fileLogger.initFile()
	return fileLogger
}

func (f *FileLogger) initFile() {
	fileAllName := fmt.Sprintf("%s/%s.log", f.filePath, f.fileName)
	file, err := os.OpenFile(fileAllName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err : %s", fileAllName, err.Error()))
	}

	f.file = file
}


func (f *FileLogger) SetLevel(level int) {
	if level < DEBUG || level > ERROR {
		level = DEBUG
	}
	f.level = level
}

func (f *FileLogger) Debug(args ...interface{}) {
	f.SetLevel(DEBUG)
	WriteLog(f.file, f.level, args...)
}

func (f *FileLogger) Info(args ...interface{}) {
	f.SetLevel(INFO)
	WriteLog(f.file, f.level, args...)
}

func (f *FileLogger) Error(args ...interface{}) {
	f.SetLevel(ERROR)
	WriteLog(f.file, f.level, args...)
}

func (f *FileLogger) Close() {
	f.file.Close()
}

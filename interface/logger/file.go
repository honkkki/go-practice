package logger

import (
	"fmt"
	"os"
	"path"
	"time"
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

// 打印日志级别与时间
func (f *FileLogger) LogPrefix() {
	now := time.Now()
	nowTime := now.Format("2006-01-02 15:04:05")
	switch f.level {
	case DEBUG:
		fmt.Fprint(f.file, "[DEBUG] ")
	case INFO:
		fmt.Fprint(f.file, "[INFO] ")
	case ERROR:
		fmt.Fprint(f.file, "[ERROR] ")
	default:
		fmt.Fprint(f.file, "[DEBUG] ")
	}
	fmt.Fprint(f.file, nowTime+" ")

	fileName, funcName, line := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	fmt.Fprintf(f.file, "%s:%d %s ", fileName, line, funcName)

}

func (f *FileLogger) SetLevel(level int) {
	if level < DEBUG || level > ERROR {
		level = DEBUG
	}
	f.level = level
}

func (f *FileLogger) Debug(args ...interface{}) {
	f.SetLevel(DEBUG)
	f.LogPrefix()
	_, err := fmt.Fprintln(f.file, args...)
	if err != nil {
		fmt.Println(err)
	}
}

func (f *FileLogger) Info(args ...interface{}) {
	f.SetLevel(INFO)
	f.LogPrefix()
	_, err := fmt.Fprintln(f.file, args...)
	if err != nil {
		fmt.Println(err)
	}
}

func (f *FileLogger) Error(args ...interface{}) {
	f.SetLevel(ERROR)
	f.LogPrefix()
	_, err := fmt.Fprintln(f.file, args...)
	if err != nil {
		fmt.Println(err)
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
}

package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}

	return
}

func WriteLog(f *os.File, LogLevel int, args ...interface{})  {
	now := time.Now()
	nowTime := now.Format("2006-01-02 15:04:05")
	switch LogLevel {
	case DEBUG:
		fmt.Fprint(f, "[DEBUG] ")
	case INFO:
		fmt.Fprint(f, "[INFO] ")
	case ERROR:
		fmt.Fprint(f, "[ERROR] ")
	default:
		fmt.Fprint(f, "[DEBUG] ")
	}
	fmt.Fprint(f, nowTime+" ")

	fileName, funcName, line := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	fmt.Fprintf(f, "%s:%d %s ", fileName, line, funcName)

	_, err := fmt.Fprintln(f, args...)
	if err != nil {
		fmt.Println(err)
	}
}

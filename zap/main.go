package main

import (
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarlogger *zap.SugaredLogger
var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("https://www.google.com")
}

func InitLogger() {
	//logger, _ = zap.NewProduction()
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
	sugarlogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	//encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./test.log")
	file, err := os.OpenFile("./test.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return zapcore.AddSync(file)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
		return
	} else {
		logger.Debug("", zap.String("url", url))
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
	}
	resp.Body.Close()
}

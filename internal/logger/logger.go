package logger

import (
	"DHT/internal/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

// Init the logger
func Init(logFilePath string, logLevel zapcore.Level) error {
	const LOG_FOLDER string = "./logs"
	if Logger != nil {
		Sync()
	}
	utils.CheckAndMakeDir(LOG_FOLDER)
	//cfg := zap.NewProductionConfig()
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{LOG_FOLDER + "/" + logFilePath}
	cfg.Level = zap.NewAtomicLevelAt(logLevel)
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if logger, err := cfg.Build(); err != nil {
		return err
	} else {
		Logger = logger.Sugar()
	}
	return nil
}

// Sync flushes any buffered log entries.
func Sync() {
	if Logger != nil {
		err := Logger.Sync()
		if err != nil {
			fmt.Println("log flush error", err)
		}
	}
}

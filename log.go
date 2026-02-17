package quietlog

import (
	"fmt"

	"github.com/varadekd/quietlog/logger"
)

func Init() {
	logger.Init("")
}

// For custom PATH
func Config(path string) {
	logger.Init(path)
}

func Info(msg string) {
	logger.Log(logger.InfoLevel, msg)
}

func Debug(msg string) {
	logger.Log(logger.DebugLevel, msg)
}

func Warn(msg string) {
	logger.Log(logger.WarnLevel, msg)
}

func Error(msg string) {
	logger.Log(logger.ErrorLevel, msg)
}

func Fatal(msg string) {
	logger.Log(logger.FatalLevel, msg)
	panic(msg)
}

func Infof(format string, args ...any) {
	logger.Logf(logger.InfoLevel, format, args...)
}

func Debugf(format string, args ...any) {
	logger.Logf(logger.DebugLevel, format, args...)
}

func Warnf(format string, args ...any) {
	logger.Logf(logger.WarnLevel, format, args...)
}

func Errorf(format string, args ...any) {
	logger.Logf(logger.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...any) {
	logger.Logf(logger.FatalLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

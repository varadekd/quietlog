package logger

import "fmt"

func logLine(lvl level, msg string) {
	Init()

	mu.Lock()
	defer mu.Unlock()

	if lines >= cfg.MaxLines {
		lines = 0
	}

	logger.Println(formatLine(lvl, msg))
	lines++
}

func Info(msg string)  { logLine(INFO, msg) }
func Debug(msg string) { logLine(DEBUG, msg) }
func Error(msg string) { logLine(ERROR, msg) }
func Fatal(msg string) {
	logLine(FATAL, msg)
	panic(msg)
}

func Infof(f string, a ...any)  { Info(fmt.Sprintf(f, a...)) }
func Debugf(f string, a ...any) { Debug(fmt.Sprintf(f, a...)) }
func Errorf(f string, a ...any) { Error(fmt.Sprintf(f, a...)) }
func Fatalf(f string, a ...any) { Fatal(fmt.Sprintf(f, a...)) }

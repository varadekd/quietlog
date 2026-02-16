package logger

import (
	"fmt"
)

func Log(level Level, msg string) {
	if !initialized.Load() {
		Init("")
	}

	// Terminating the debug logs
	if level == DebugLevel && !cfg.DebugLevel {
		return
	}

	if lineCnt.Add(1) > cfg.MaxLines {
		lineCnt.Store(1)
	}

	base.Println(formatLine(level, msg))
}

func Logf(level Level, format string, args ...any) {
	Log(level, fmt.Sprintf(format, args...))
}

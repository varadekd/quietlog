package logger

import (
	"fmt"
)

func Log(level Level, msg string) {
	if !initialized.Load() {
		Init()
	}

	if lineCnt.Add(1) > cfg.MaxLines {
		lineCnt.Store(1)
	}

	base.Println(formatLine(level, msg))
}

func Logf(level Level, format string, args ...any) {
	Log(level, fmt.Sprintf(format, args...))
}

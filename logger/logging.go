package logger

import (
	"fmt"
)

func Log(level Level, msg string) {
	Init("")

	if level == DebugLevel && !cfg.DebugLevel {
		return
	}

	line := formatLine(level, msg)
	base.Println(line)

	// track size (approximate)
	currentSize.Add(int64(len(line) + 1)) // +1 for newline

	// rotate if over threshold
	if currentSize.Load() > cfg.MaxSizeInMb*1024*1024 {
		rotate()
	}
}

func Logf(level Level, format string, args ...any) {
	Log(level, fmt.Sprintf(format, args...))
}

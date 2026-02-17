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

	// Check if rotation needed BEFORE writing
	lineSize := int64(len(line) + 1) // +1 for newline
	maxBytes := int64(cfg.MaxSizeInMb * 1024 * 1024)
	if currentSize.Load()+lineSize > maxBytes {
		rotate()
	}

	base.Println(line)
	currentSize.Add(lineSize)
}

func Logf(level Level, format string, args ...any) {
	Log(level, fmt.Sprintf(format, args...))
}

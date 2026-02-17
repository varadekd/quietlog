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
	lineSize := int64(len(line) + 1)

	rotateMu.Lock()
	maxBytes := int64(cfg.MaxSizeInMb * 1024 * 1024)
	if maxBytes <= 0 {
		maxBytes = 10 * 1024 * 1024 // 10MB safety fallback
	}
	if currentSize.Load()+lineSize > maxBytes {
		rotateLocked()
	}
	base.Println(line)
	currentSize.Add(lineSize)
	rotateMu.Unlock()
}

func Logf(level Level, format string, args ...any) {
	Log(level, fmt.Sprintf(format, args...))
}

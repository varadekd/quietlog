package logger

import (
	"fmt"
	"os"
)

func Log(level Level, msg string) {
	Init("")

	if level == DebugLevel && !cfg.DebugLevel {
		return
	}

	plainLine := formatLine(level, msg, false)
	colorLine := formatLine(level, msg, cfg.Color)

	lineSize := int64(len(plainLine) + 1)

	rotateMu.Lock()
	maxBytes := int64(cfg.MaxSizeInMb * 1024 * 1024)
	if maxBytes <= 0 {
		maxBytes = 10 * 1024 * 1024
	}
	if currentSize.Load()+lineSize > maxBytes {
		rotateLocked()
	}

	// stdout gets color, file gets plain text
	if cfg.FileLogging && currentFile != nil {
		fmt.Fprintln(os.Stdout, colorLine)
		fmt.Fprintln(currentFile, plainLine)
	} else {
		fmt.Fprintln(os.Stdout, colorLine)
	}

	currentSize.Add(lineSize)
	rotateMu.Unlock()
}

func Logf(level Level, format string, args ...any) {
	Log(level, fmt.Sprintf(format, args...))
}

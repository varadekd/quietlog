package logger

import (
	"fmt"
	"sync/atomic"
)

func Log(level Level, msg string) {
	Init()

	if atomic.AddInt64(&lineCnt, 1) > cfg.MaxLines {
		atomic.StoreInt64(&lineCnt, 1)
	}

	base.Println(formatLine(level, msg))
}

func Logf(level Level, format string, args ...any) {
	Log(level, fmt.Sprintf(format, args...))
}

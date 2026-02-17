package logger

import (
	"log"
	"os"
	"sync"
	"sync/atomic"
)

type Level string

const (
	InfoLevel  Level = "INFO"
	DebugLevel Level = "DEBUG"
	WarnLevel  Level = "WARN"
	ErrorLevel Level = "ERROR"
	FatalLevel Level = "FATAL"
)

var (
	once         sync.Once
	rotateMu     sync.Mutex
	base         *log.Logger
	cfg          Config
	currentFile  *os.File
	currentSize  atomic.Int64
	currentChunk atomic.Int32
	startTime    string
)

// ResetForTesting resets the logger state - ONLY for tests
func ResetForTesting() {
	rotateMu.Lock()
	defer rotateMu.Unlock()

	if currentFile != nil && currentFile != os.Stdout {
		currentFile.Close()
	}

	once = sync.Once{}
	currentFile = nil
	currentSize.Store(0)
	currentChunk.Store(0)
	startTime = ""
	base = nil
}

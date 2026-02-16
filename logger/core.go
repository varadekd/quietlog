package logger

import (
	"log"
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
	mu          sync.Mutex
	initialized atomic.Bool
	base        *log.Logger
	cfg         Config
	lineCnt     atomic.Int64
)

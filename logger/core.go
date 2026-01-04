package logger

import (
	"log"
	"sync"
)

type Level string

const (
	InfoLevel  Level = "INFO"
	DebugLevel Level = "DEBUG"
	ErrorLevel Level = "ERROR"
	FatalLevel Level = "FATAL"
)

var (
	once    sync.Once
	base    *log.Logger
	lineCnt int64
	cfg     Config
)

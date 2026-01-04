package logger

import (
	"log"
	"sync"
	"time"
)

type level string

const (
	INFO  level = "INFO"
	DEBUG level = "DEBUG"
	ERROR level = "ERROR"
	FATAL level = "FATAL"
)

var (
	once   sync.Once
	logger *log.Logger
	cfg    config
	mu     sync.Mutex
	lines  int
)

type config struct {
	AppName    string
	Color      bool
	FilePath   string
	MaxLines   int
	Format     []string
	TimeFormat string
	Location   *time.Location
}

func Init() {
	once.Do(func() {
		loadConfig()
		writer := buildWriter()
		logger = log.New(writer, "", 0)
	})
}

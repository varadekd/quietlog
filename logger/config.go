package logger

import "time"

type Config struct {
	AppName    string
	FilePath   string
	MaxLines   int64
	Format     []string
	TimeFmt    string
	Location   *time.Location
	Color      bool
	DebugLevel bool
}

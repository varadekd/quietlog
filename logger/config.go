package logger

import "time"

type Config struct {
	AppName     string
	FilePath    string
	MaxSizeInMb float64
	CleanLogs   bool
	Format      []string
	TimeFmt     string
	Location    *time.Location
	Color       bool
	DebugLevel  bool
}

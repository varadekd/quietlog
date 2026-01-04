package logger

import (
	"io"
	"os"
)

func buildWriter(cfg Config) io.Writer {
	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return os.Stdout
	}
	return io.MultiWriter(os.Stdout, file)
}

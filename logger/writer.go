package logger

import (
	"io"
	"os"
)

func buildWriter() io.Writer {
	file, err := os.OpenFile(cfg.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return os.Stdout
	}
	return io.MultiWriter(os.Stdout, file)
}

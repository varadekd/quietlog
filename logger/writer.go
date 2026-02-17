package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func buildWriter(cfg Config) io.Writer {
	// Ensure directory exists
	dir := filepath.Dir(cfg.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "quietlog: cannot create log directory %s: %v\n", dir, err)
		return os.Stdout
	}

	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "quietlog: cannot open log file %s: %v\n", cfg.FilePath, err)
		return os.Stdout
	}

	return io.MultiWriter(os.Stdout, file)
}

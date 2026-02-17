package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Init(filepath string) {
	once.Do(func() {
		cfg = LoadConfig(filepath)

		if cfg.FileLogging {
			startTime = time.Now().Format("20060102_150405")
			currentChunk.Store(0)

			if cfg.CleanLogs {
				cleanOldLogs(cfg.FilePath, cfg.AppName)
			}
		}

		rotateLocked()
	})
}

func cleanOldLogs(dir, appName string) {
	pattern := filepath.Join(dir, appName+"_*_chunk*.log")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "quietlog: failed to list old logs: %v\n", err)
		return
	}
	for _, f := range matches {
		if err := os.Remove(f); err != nil {
			fmt.Fprintf(os.Stderr, "quietlog: failed to remove %s: %v\n", f, err)
		}
	}
}

// rotateLocked does the actual rotation work.
// Caller must already hold rotateMu — OR be inside once.Do (init time).
func rotateLocked() {
	if currentFile != nil && currentFile != os.Stdout {
		currentFile.Close()
		currentFile = nil
	}

	var writer io.Writer = os.Stdout

	if cfg.FileLogging {
		chunk := currentChunk.Add(1)
		filename := fmt.Sprintf("%s_%s_chunk%03d.log",
			cfg.AppName, startTime, chunk)
		logPath := filepath.Join(cfg.FilePath, filename)

		f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "quietlog: failed to open log file %s: %v\n", logPath, err)
			// stdout still works, just no file
		} else {
			currentFile = f
			writer = io.MultiWriter(os.Stdout, f)
		}
	}

	currentSize.Store(0)
	base = log.New(writer, "", 0)
}

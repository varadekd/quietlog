package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Init(filepath string) {
	once.Do(func() {
		cfg = LoadConfig(filepath)

		// clean old logs if enabled
		if cfg.CleanLogs {
			cleanOldLogs(cfg.FilePath, cfg.AppName)
		}

		// set shared timestamp for all chunks
		startTime = time.Now().Format("20060102_150405")
		currentChunk.Store(0)

		rotate() // open first chunk
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

func rotate() {
	rotateMu.Lock()
	defer rotateMu.Unlock()

	// close old file
	if currentFile != nil {
		currentFile.Close()
	}

	// open new chunk
	chunk := currentChunk.Add(1)
	filename := fmt.Sprintf("%s_%s_chunk%03d.log",
		cfg.AppName, startTime, chunk)

	logPath := filepath.Join(cfg.FilePath, filename)

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "quietlog: failed to open log file %s: %v\n", logPath, err)
		// fallback to stdout
		f = os.Stdout
	}

	currentFile = f
	currentSize.Store(0)
	base = log.New(f, "", 0)
}

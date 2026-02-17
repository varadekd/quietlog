package quietlog

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/varadekd/quietlog/logger"
)

func TestConcurrentPublicLogging(t *testing.T) {
	logger.ResetForTesting()

	tmpDir := t.TempDir()
	currentConfig := "quietlog_config.json"
	os.WriteFile(currentConfig, []byte(`{
		"app_name": "testapp",
		"file_logging": true,
		"log_file_path": "`+tmpDir+`"
	}`), 0644)
	defer os.Remove(currentConfig)

	Init()

	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Info("info log")
			Warn("warning log")
			Debug("debug log")
			Error("error log")
		}(i)
	}
	wg.Wait()

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("failed to read log directory: %v", err)
	}

	var logFiles int
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".log") {
			logFiles++
		}
	}
	if logFiles == 0 {
		t.Fatal("no log files created")
	}
}

func TestFormattedLogging(t *testing.T) {
	logger.ResetForTesting()

	tmpDir := t.TempDir()
	currentConfig := "quietlog_config.json"
	os.WriteFile(currentConfig, []byte(`{
		"app_name": "testapp",
		"file_logging": true,
		"log_file_path": "`+tmpDir+`"
	}`), 0644)
	defer os.Remove(currentConfig)

	Init()
	Infof("user %s logged in from %s", "alice", "192.168.1.1")
	Debugf("cache hit rate: %.2f%%", 95.67)
	Warnf("disk usage: %d%%", 85)
	Errorf("failed to connect to %s:%d", "db.example.com", 5432)

	entries, _ := os.ReadDir(tmpDir)
	var logFile string
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".log") {
			logFile = filepath.Join(tmpDir, e.Name())
			break
		}
	}
	if logFile == "" {
		t.Fatal("no log file created")
	}

	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	logs := string(content)
	tests := []string{
		"user alice logged in from 192.168.1.1",
		"disk usage: 85%",
		"failed to connect to db.example.com:5432",
	}
	for _, want := range tests {
		if !strings.Contains(logs, want) {
			t.Errorf("log file missing formatted message: %q", want)
		}
	}
}

package logger

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func TestInternalConcurrentLogging(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()

	// Write config file
	configPath := filepath.Join(tmpDir, "quietlog_config.json")
	os.WriteFile(configPath, []byte(`{
		"app_name": "testapp",
		"log_file_path": "`+tmpDir+`",
		"max_file_size_mb": 10,
		"clean_logs": false
	}`), 0644)

	Init(configPath)

	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Log(InfoLevel, "internal log")
		}()
	}
	wg.Wait()

	// Verify logs were written
	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("failed to read log dir: %v", err)
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

func TestInvalidLogFilePathDoesNotPanic(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "quietlog_config.json")

	// Config with invalid path
	os.WriteFile(configPath, []byte(`{
		"app_name": "testapp",
		"log_file_path": "/invalid/path/nowhere",
		"max_file_size_mb": 10
	}`), 0644)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("logger panicked on invalid file path: %v", r)
		}
	}()

	Init(configPath)
	Log(InfoLevel, "should fallback to stdout")
}

func TestDebugLevelToggle(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "quietlog_config.json")

	// Debug disabled
	os.WriteFile(configPath, []byte(`{
		"app_name": "testapp",
		"log_file_path": "`+tmpDir+`",
		"debug_level": false
	}`), 0644)

	Init(configPath)
	Log(DebugLevel, "debug message")
	Log(InfoLevel, "info message")

	// Read log file
	entries, _ := os.ReadDir(tmpDir)
	var logFile string
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".log") {
			logFile = filepath.Join(tmpDir, e.Name())
			break
		}
	}

	content, _ := os.ReadFile(logFile)
	logs := string(content)

	if strings.Contains(logs, "debug message") {
		t.Error("debug message appeared when debug_level was false")
	}
	if !strings.Contains(logs, "info message") {
		t.Error("info message missing when debug_level was false")
	}
}

func TestFileRotation(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "quietlog_config.json")

	// Very small file size to trigger rotation - 0.0005 MB = ~512 bytes
	os.WriteFile(configPath, []byte(`{
			"app_name": "testapp",
			"log_file_path": "`+tmpDir+`",
			"max_file_size_mb": 0.0005
		}`), 0644)

	Init(configPath)

	// Write enough logs to trigger rotation - each line is ~80 bytes
	for i := 0; i < 200; i++ {
		Log(InfoLevel, "this is a log message that will trigger rotation when accumulated")
	}

	// Check for multiple chunk files
	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("failed to read log dir: %v", err)
	}

	var chunkFiles []string
	for _, e := range entries {
		if strings.Contains(e.Name(), "_chunk") {
			chunkFiles = append(chunkFiles, e.Name())
		}
	}

	if len(chunkFiles) < 2 {
		t.Errorf("expected at least 2 chunk files, got %d", len(chunkFiles))
	}
}

func TestCleanLogsFlag(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()

	// Create some old log files
	os.WriteFile(filepath.Join(tmpDir, "testapp_20240101_120000_chunk001.log"), []byte("old"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "testapp_20240101_120000_chunk002.log"), []byte("old"), 0644)

	configPath := filepath.Join(tmpDir, "quietlog_config.json")
	os.WriteFile(configPath, []byte(`{
		"app_name": "testapp",
		"log_file_path": "`+tmpDir+`",
		"clean_logs": true
	}`), 0644)

	Init(configPath)
	Log(InfoLevel, "new log")

	// Old files should be gone
	entries, _ := os.ReadDir(tmpDir)
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), "testapp_20240101") {
			t.Errorf("old log file was not cleaned: %s", e.Name())
		}
	}
}

func TestInvalidTimezone(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "quietlog_config.json")

	os.WriteFile(configPath, []byte(`{
		"app_name": "testapp",
		"log_file_path": "`+tmpDir+`",
		"log_timezone": "Invalid/Timezone"
	}`), 0644)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("logger panicked on invalid timezone: %v", r)
		}
	}()

	Init(configPath)
	Log(InfoLevel, "should use UTC fallback")
}

func TestInvalidJSON(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "quietlog_config.json")

	// Malformed JSON
	os.WriteFile(configPath, []byte(`{
		"app_name": "testapp"
		"log_file_path": "`+tmpDir+`"
	}`), 0644)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("logger panicked on invalid JSON: %v", r)
		}
	}()

	Init(configPath)
	Log(InfoLevel, "should use defaults")
}

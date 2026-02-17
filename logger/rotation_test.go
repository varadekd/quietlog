package logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestChunkNaming(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "quietlog_config.json")

	os.WriteFile(configPath, []byte(`{
		"app_name": "myapp",
		"log_file_path": "`+tmpDir+`",
		"max_file_size_mb": 0.001
	}`), 0644)

	Init(configPath)

	// Write enough to create multiple chunks
	for i := 0; i < 100; i++ {
		Log(InfoLevel, "log message to trigger chunk rotation")
	}

	entries, _ := os.ReadDir(tmpDir)
	var chunks []string
	for _, e := range entries {
		if strings.Contains(e.Name(), "_chunk") {
			chunks = append(chunks, e.Name())
		}
	}

	// Verify chunk naming format
	for _, chunk := range chunks {
		if !strings.HasPrefix(chunk, "myapp_") {
			t.Errorf("chunk name doesn't start with app name: %s", chunk)
		}
		if !strings.Contains(chunk, "_chunk") {
			t.Errorf("chunk name missing '_chunk': %s", chunk)
		}
		if !strings.HasSuffix(chunk, ".log") {
			t.Errorf("chunk name doesn't end with .log: %s", chunk)
		}
	}
}

func TestRotationPreservesAllLogs(t *testing.T) {
	ResetForTesting()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "quietlog_config.json")

	os.WriteFile(configPath, []byte(`{
		"app_name": "testapp",
		"log_file_path": "`+tmpDir+`",
		"max_file_size_mb": 0.001
	}`), 0644)

	Init(configPath)

	totalLogs := 200
	for i := 0; i < totalLogs; i++ {
		Log(InfoLevel, "message number")
	}

	// Count total log lines across all chunks
	entries, _ := os.ReadDir(tmpDir)
	totalLines := 0
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".log") {
			content, _ := os.ReadFile(filepath.Join(tmpDir, e.Name()))
			totalLines += strings.Count(string(content), "\n")
		}
	}

	if totalLines < totalLogs {
		t.Errorf("expected at least %d log lines, got %d (logs lost during rotation)", totalLogs, totalLines)
	}
}

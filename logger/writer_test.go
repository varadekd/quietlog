package logger

import "testing"

func TestInvalidLogFilePathDoesNotPanic(t *testing.T) {
	t.Setenv("LOG_FILE_PATH", "/invalid/path/nowhere.log")

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("logger panicked on invalid file path: %v", r)
		}
	}()

	Init()
	Log(InfoLevel, "should fallback to stdout")
}

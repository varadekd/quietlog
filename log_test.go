package quietlog

import "testing"

func TestBasicLoggingDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("logging panicked: %v", r)
		}
	}()

	Info("info message")
	Debug("debug message")
	Error("error message")
}

func TestEnvConfigDoesNotPanic(t *testing.T) {
	t.Setenv("APP_NAME", "testapp")
	t.Setenv("LOG_LINES", "10")
	t.Setenv("LOG_FORMAT", "Timestamp,Message")

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("env-based logging panicked: %v", r)
		}
	}()

	Info("hello from env test")
}

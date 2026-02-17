package main

import (
	"time"

	"github.com/varadekd/quietlog"
)

func main() {
	// --------------------------------------------------
	// 1. Zero config — just works
	// --------------------------------------------------
	quietlog.Info("quietlog started — no config needed")

	// --------------------------------------------------
	// 2. All log levels
	// --------------------------------------------------
	quietlog.Info("user alice logged in")
	quietlog.Debug("cache miss for key: session_abc123") // silent by default
	quietlog.Warn("disk usage at 82% — consider cleanup")
	quietlog.Error("failed to send email: connection refused")

	// --------------------------------------------------
	// 3. Formatted variants
	// --------------------------------------------------
	quietlog.Infof("server listening on port %d", 8080)
	quietlog.Warnf("response time %.2fms exceeds threshold", 312.5)
	quietlog.Errorf("database query failed after %d retries", 3)

	// --------------------------------------------------
	// 4. Simulate a real service startup sequence
	// --------------------------------------------------
	quietlog.Info("loading configuration")
	time.Sleep(10 * time.Millisecond)

	quietlog.Info("connecting to database")
	time.Sleep(10 * time.Millisecond)

	quietlog.Info("running migrations")
	time.Sleep(10 * time.Millisecond)

	quietlog.Warn("migration 0042 skipped — already applied")
	time.Sleep(10 * time.Millisecond)

	quietlog.Info("starting HTTP server on :8080")
	quietlog.Info("ready to serve requests")

	// --------------------------------------------------
	// 5. Fatal — logs then panics (recovered here for demo)
	// --------------------------------------------------
	func() {
		defer func() {
			if r := recover(); r != nil {
				quietlog.Errorf("recovered from fatal: %v", r)
			}
		}()
		quietlog.Fatal("could not bind to port 8080: address already in use")
	}()

	quietlog.Info("example complete")
}

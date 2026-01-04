package logger

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var envOnce sync.Once

func ensureEnvLoaded() {
	envOnce.Do(func() {
		// Heuristic: if APP_NAME is already set, assume env is loaded
		if os.Getenv("APP_NAME") != "" {
			return
		}

		// Best-effort load .env from working directory
		_ = godotenv.Load()
	})
}

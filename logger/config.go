package logger

import (
	"os"
	"time"
)

func loadConfig() {
	cfg.AppName = getEnv("APP_NAME", "app")
	cfg.Color = getEnvBool("LOG_COLOR", true)
	cfg.MaxLines = getEnvInt("LOG_LINES", 1000)
	cfg.TimeFormat = getEnv("LOG_TIME_FORMAT", "02-Jan-2006 15:04:05")

	format := getEnv("LOG_FORMAT", "AppName,Level,Timestamp,Message")
	cfg.Format = splitAndTrim(format)

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		loc = time.Local
	}
	cfg.Location = loc

	cfg.FilePath = os.Getenv("LOG_FILE_PATH")
	if cfg.FilePath == "" {
		ts := time.Now().Format("20060102_150405")
		cfg.FilePath = "./" + cfg.AppName + "_" + ts + ".log"
	}
}

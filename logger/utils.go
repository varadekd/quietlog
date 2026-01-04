package logger

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func loadConfig() Config {
	app := getEnv("APP_NAME", "app")

	path := os.Getenv("LOG_FILE_PATH")
	if path == "" {
		path = "./" + app + "_" + time.Now().Format("20060102_150405") + ".log"
	}

	format := getEnv("LOG_FORMAT", "AppName,Level,Timestamp,Message")

	loc, _ := time.LoadLocation("Asia/Kolkata")

	return Config{
		AppName:  app,
		FilePath: path,
		MaxLines: int64(getEnvInt("LOG_LINES", 1000)),
		Format:   split(format),
		TimeFmt:  getEnv("LOG_TIME_FORMAT", "02-Jan-2006 15:04:05"),
		Location: loc,
	}
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func getEnvInt(k string, d int) int {
	if v := os.Getenv(k); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return d
}

func split(s string) []string {
	p := strings.Split(s, ",")
	for i := range p {
		p[i] = strings.TrimSpace(p[i])
	}
	return p
}

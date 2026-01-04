package logger

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func loadConfig() Config {
	ensureEnvLoaded() // 👈 FAIL-SAFE

	app := getEnv("APP_NAME", "app")

	path := os.Getenv("LOG_FILE_PATH")
	if path == "" {
		path = "./" + app + "_" + time.Now().Format("20060102_150405") + ".log"
	}

	timezone := getEnv("LOG_TIMEZONE", "Asia/Kolkata")

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.Local
	}

	return Config{
		AppName:  app,
		FilePath: path,
		MaxLines: int64(getEnvInt("LOG_LINES", 1000)),
		Format:   split(getEnv("LOG_FORMAT", "AppName,Level,Timestamp,Message")),
		TimeFmt:  getEnv("LOG_TIME_FORMAT", "02-Jan-2006 15:04:05"),
		Location: loc,
		Color:    getEnvBool("LOG_COLOR", true),
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

func getEnvBool(k string, d bool) bool {
	v := strings.ToLower(os.Getenv(k))
	if v == "true" || v == "1" {
		return true
	}
	if v == "false" || v == "0" {
		return false
	}
	return d
}

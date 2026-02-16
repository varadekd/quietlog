package logger

import (
	"encoding/json"
	"os"
	"time"
)

/*
jsonConfig mirrors the JSON structure.
*/
type jsonConfig struct {
	AppName    string   `json:"app_name"`
	FilePath   string   `json:"log_file_path"`
	Timezone   string   `json:"log_timezone"`
	MaxLines   int64    `json:"log_lines"`
	Format     []string `json:"log_format"`
	TimeFmt    string   `json:"log_time_format"`
	Color      bool     `json:"log_color"`
	DebugLevel bool     `json:"debug_level"`
	Quiet      bool     `json:"quiet"`
}

/*
loadConfig reads configuration from logger.json.
If file is missing or invalid → defaults are used.

System NEVER crashes.
*/
func loadConfig() Config {
	const path = "./quietlog_config.json"

	// defaults
	jc := jsonConfig{
		AppName:    "",
		FilePath:   "",
		Timezone:   "Asia/Kolkata",
		MaxLines:   1000,
		Format:     []string{"AppName", "Level", "Timestamp", "Message"},
		TimeFmt:    "02-Jan-2006 15:04:05",
		Color:      true,
		DebugLevel: false,
	}

	// read file if present
	b, err := os.ReadFile(path)
	if err != nil {
		Logf(WarnLevel, "logger config not found at %s — using defaults", path)
	} else {
		if err := json.Unmarshal(b, &jc); err != nil {
			Logf(WarnLevel, "logger config invalid JSON at %s — using defaults: %v", path, err)
		}
	}

	// auto file name if empty
	if jc.FilePath == "" {
		jc.FilePath = "./" + jc.AppName + "_" + time.Now().Format("20060102_150405") + ".log"
	}

	// timezone handling
	loc, err := time.LoadLocation(jc.Timezone)
	if err != nil {
		Logf(WarnLevel, "invalid timezone %q — falling back to system local", jc.Timezone)
		loc = time.Local
	}

	return Config{
		AppName:    jc.AppName,
		FilePath:   jc.FilePath,
		MaxLines:   jc.MaxLines,
		Format:     jc.Format,
		TimeFmt:    jc.TimeFmt,
		Location:   loc,
		Color:      jc.Color,
		DebugLevel: jc.DebugLevel,
	}
}

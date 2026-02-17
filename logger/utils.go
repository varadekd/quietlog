package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

/*
jsonConfig mirrors the Config structure.
*/
type jsonConfig struct {
	AppName     string   `json:"app_name"`
	FileLogging bool     `json:"file_logging"`
	FilePath    string   `json:"log_file_path"`
	MaxSizeInMb float64  `json:"max_file_size_mb"`
	CleanLogs   bool     `json:"clean_logs"`
	Timezone    string   `json:"log_timezone"`
	Format      []string `json:"log_format"`
	TimeFmt     string   `json:"log_time_format"`
	Color       bool     `json:"log_color"`
	DebugLevel  bool     `json:"debug_level"`
}

/*
loadConfig reads configuration from logger.json.
If file is missing or invalid → defaults are used.
System NEVER crashes.
*/
func LoadConfig(configFile string) Config {
	path := "./quietlog_config.json"

	if configFile != "" {
		path = configFile
	}

	// defaults
	jc := jsonConfig{
		AppName:     "APP",
		FileLogging: false,
		FilePath:    "",
		Timezone:    "UTC",
		MaxSizeInMb: 10,
		CleanLogs:   false,
		Format:      []string{"AppName", "Level", "Message", "Timestamp"},
		TimeFmt:     "02-Jan-2006 15:04:05 MST",
		Color:       false,
		DebugLevel:  false,
	}

	// read file if present
	b, err := os.ReadFile(path)
	if err != nil {
		if configFile != "" {
			fmt.Printf("quietlog: config not found at %s — using defaults\n", path)
		}
	} else {
		if err := json.Unmarshal(b, &jc); err != nil {
			fmt.Printf("logger config invalid JSON at %s — using defaults: %v\n", path, err)
		}
	}

	if jc.FileLogging && jc.FilePath == "" {
		jc.FilePath = "./"
	}

	if !jc.FileLogging && jc.FilePath != "" {
		fmt.Printf("quietlog: log_file_path is set but file_logging is false — no file will be created\n")
	}

	// timezone handling
	loc, err := time.LoadLocation(jc.Timezone)
	if err != nil {
		fmt.Printf("invalid timezone %q — falling back to UTC\n", jc.Timezone)
		loc = time.UTC
	}

	return Config{
		AppName:     jc.AppName,
		FileLogging: jc.FileLogging,
		FilePath:    jc.FilePath,
		MaxSizeInMb: jc.MaxSizeInMb,
		CleanLogs:   jc.CleanLogs,
		Format:      jc.Format,
		TimeFmt:     jc.TimeFmt,
		Location:    loc,
		Color:       jc.Color,
		DebugLevel:  jc.DebugLevel,
	}
}

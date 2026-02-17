package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
func LoadConfig(configFile string) Config {
	path := "quietlog_config.json"

	if configFile != "" {
		path = configFile
	}

	// defaults
	jc := jsonConfig{
		AppName:    "APP",
		FilePath:   "",
		Timezone:   "",
		MaxLines:   1000,
		Format:     []string{"AppName", "Level", "Message", "Timestamp"},
		TimeFmt:    "02-Jan-2006 15:04:05 MST",
		Color:      true,
		DebugLevel: false,
	}

	// read file if present
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("logger config not found at %s — using defaults\n", path)
	} else {
		if err := json.Unmarshal(b, &jc); err != nil {
			fmt.Printf("logger config invalid JSON at %s — using defaults: %v\n", path, err)
		}
	}

	fmt.Printf("Loaded config %+v \n", jc)

	// auto file name if empty

	filename := jc.AppName + "_" + time.Now().Format("20060102_150405") + ".log"

	switch {
	case jc.FilePath == "":
		// default → current directory
		jc.FilePath = filename

	case filepath.Ext(jc.FilePath) == "":
		// path is a directory
		jc.FilePath = filepath.Join(jc.FilePath, filename)

	default:
		// user provided a file → respect it
		// do nothing
	}

	// timezone handling
	loc, err := time.LoadLocation(jc.Timezone)
	if err != nil {
		fmt.Printf("invalid timezone %q — falling back to system local", jc.Timezone)
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

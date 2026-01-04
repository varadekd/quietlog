package logger

import (
	"strings"
	"time"
)

func formatLine(lvl level, msg string) string {
	now := time.Now().In(cfg.Location).Format(cfg.TimeFormat)

	parts := map[string]string{
		"AppName":   cfg.AppName,
		"Level":     string(lvl),
		"Timestamp": now,
		"Message":   msg,
	}

	var out []string
	for _, f := range cfg.Format {
		if v, ok := parts[f]; ok {
			out = append(out, "["+v+"]")
		}
	}
	return strings.Join(out, " ") + " : " + msg
}

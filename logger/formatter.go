package logger

import (
	"strings"
	"time"
)

func formatLine(level Level, msg string) string {
	now := time.Now().In(cfg.Location).Format(cfg.TimeFmt)

	fields := map[string]string{
		"AppName":   cfg.AppName,
		"Level":     string(level),
		"Timestamp": now,
		"Message":   msg,
	}

	var b strings.Builder
	for _, f := range cfg.Format {
		if v, ok := fields[f]; ok {
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("] ")
		}
	}

	b.WriteString(": ")
	b.WriteString(msg)

	return b.String()
}

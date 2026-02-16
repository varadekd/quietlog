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

	for i, f := range cfg.Format {
		if v, ok := fields[f]; ok {
			if i > 0 {
				b.WriteString(" ")
			}

			// Apply color ONLY to Level (and optionally AppName)
			if f == "Level" {
				v = colorize(level, v)
			}

			if f == "Message" {
				b.WriteString(": ")
				b.WriteString(v)
			} else {
				if v != "" {
					b.WriteString("[")
					b.WriteString(v)
					b.WriteString("]")
				}
			}
		}
	}

	return b.String()
}

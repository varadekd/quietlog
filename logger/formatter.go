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
	written := 0

	for _, f := range cfg.Format {
		v, ok := fields[f]
		if !ok {
			continue
		}

		if f == "Message" {
			if written > 0 {
				b.WriteString(" ")
			}
			b.WriteString(": ")
			b.WriteString(v)
			written++
		} else {
			if v == "" {
				continue // skip empty fields entirely
			}
			if written > 0 {
				b.WriteString(" ")
			}
			if f == "Level" {
				v = colorize(level, v)
			}
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			written++
		}
	}

	return b.String()
}

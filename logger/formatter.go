package logger

import (
	"strings"
	"time"
)

func formatLine(level Level, msg string) string {
	loc := cfg.Location
	if loc == nil {
		loc = time.UTC
	}
	timeFmt := cfg.TimeFmt
	if timeFmt == "" {
		timeFmt = "02-Jan-2006 15:04:05 MST"
	}

	now := time.Now().In(loc).Format(timeFmt)

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
				continue
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

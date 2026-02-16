package logger

import "github.com/fatih/color"

func colorize(level Level, value string) string {
	if !cfg.Color {
		return value
	}

	switch level {
	case InfoLevel:
		return color.GreenString(value)
	case DebugLevel:
		return color.CyanString(value)
	case WarnLevel:
		return color.YellowString(value)
	case ErrorLevel:
		return color.RedString(value)
	case FatalLevel:
		return color.MagentaString(value)
	default:
		return value
	}
}

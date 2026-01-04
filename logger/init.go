package logger

import (
	"log"
)

func Init() {
	once.Do(func() {
		cfg = loadConfig()
		writer := buildWriter(cfg)
		base = log.New(writer, "", 0)
	})
}

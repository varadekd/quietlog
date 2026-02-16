package logger

import (
	"log"
)

func Init(filepath string) {
	once.Do(func() {
		cfg = LoadConfig(filepath)
		writer := buildWriter(cfg)
		base = log.New(writer, "", 0)
		lineCnt.Store(0)
		initialized.Store(true)
	})
}

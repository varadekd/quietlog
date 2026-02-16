package logger

import "log"

func Init(filepath string) {
	mu.Lock()
	defer mu.Unlock()

	cfg = LoadConfig(filepath)
	writer := buildWriter(cfg)
	base = log.New(writer, "", 0)

	lineCnt.Store(0)
	initialized.Store(true)
}

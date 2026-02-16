package logger

import "log"

func Init() {
	mu.Lock()
	defer mu.Unlock()

	cfg = LoadConfig("")
	writer := buildWriter(cfg)
	base = log.New(writer, "", 0)

	lineCnt.Store(0)
	initialized.Store(true)
}

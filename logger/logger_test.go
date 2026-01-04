package logger

import (
	"sync"
	"testing"
)

func TestConcurrentLogging(t *testing.T) {
	Init()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Log(InfoLevel, "concurrent log")
		}()
	}

	wg.Wait()
}

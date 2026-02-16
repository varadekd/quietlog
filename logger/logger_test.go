package logger

import (
	"sync"
	"testing"
)

func TestInternalConcurrentLogging(t *testing.T) {
	Init("")

	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Log(InfoLevel, "internal log")
		}()
	}

	wg.Wait()
}

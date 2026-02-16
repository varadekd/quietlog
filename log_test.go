package quietlog

import (
	"sync"
	"testing"
)

func TestConcurrentPublicLogging(t *testing.T) {
	Init()

	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Info("info log")
			Warn("Warning log")
			Debug("debug log")
			Error("error log")
		}(i)
	}

	wg.Wait()
}

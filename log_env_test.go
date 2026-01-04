package quietlog

import (
	"sync"
	"testing"
)

func TestEnvReloadDuringRuntime(t *testing.T) {
	Init()

	t.Setenv("APP_NAME", "first")
	Init()
	Info("first app")

	t.Setenv("APP_NAME", "second")
	Init()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Info("log after reload")
		}()
	}

	wg.Wait()
}

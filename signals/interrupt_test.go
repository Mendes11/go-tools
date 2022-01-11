package signals

import (
	"os"
	"sync"
	"testing"
)

func TestCallOnSignal(t *testing.T) {
	ch := make(chan os.Signal)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go callOnSignal(ch, wg.Done)
	ch <- os.Interrupt
	// Block forever if test fails
	wg.Wait()
}

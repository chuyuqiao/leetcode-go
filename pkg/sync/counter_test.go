package sync

import (
	"fmt"
	"testing"
	"time"
)

func TestCounter_WithSleep(t *testing.T) {
	counter := &AtomicCounter{}
	for i := 0; i < 50; i++ {
		go func() {
			counter.Inc()
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("Count:", counter.Load())
}

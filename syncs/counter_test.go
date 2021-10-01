package syncs

import (
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safety concurrently", func(t *testing.T) {
		wantedCounter := 1000
		counter := NewCounter()

		var wg sync.WaitGroup

		wg.Add(wantedCounter)

		for i := 1; i <= wantedCounter; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, wantedCounter)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}

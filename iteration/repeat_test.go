package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat with times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat with times is 0", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat with times is 1", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat with times is less than 0", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := ""

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Long")

	got := buffer.String()
	want := "Hello, Long"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
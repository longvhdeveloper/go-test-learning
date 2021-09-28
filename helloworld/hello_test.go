package helloworld

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Long", "")
		want := "hello, Long"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Long", "spanish")
		want := "hola, Long"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in France", func(t *testing.T) {
		got := Hello("Long", "france")
		want := "bonjour, Long"
		assertCorrectMessage(t, got, want)
	})
}

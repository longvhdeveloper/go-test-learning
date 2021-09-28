package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q given %s", got, want, "test")
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinitions(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", word)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func TestSearch(t *testing.T) {
	t.Run("know word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, err := dictionary.Search("test")
		want := "this is just a test"

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		assertErrors(t, err, ErrNotFound)
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		assertErrors(t, err, nil)
		assertDefinitions(t, dictionary, "test", "this is just a test")
	})

	t.Run("exist word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Add("test", "new test value")
		assertErrors(t, err, ErrWordExist)
		assertDefinitions(t, dictionary, "test", "this is just a test")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word exist", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Update("test", "new test value")
		assertErrors(t, err, nil)
		assertDefinitions(t, dictionary, "test", "new test value")
	})

	t.Run("word not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "new test value")
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("word exist", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Delete("test")
		assertErrors(t, err, nil)

		_, searchErr := dictionary.Search("test")
		assertErrors(t, searchErr, ErrNotFound)
	})

	t.Run("word not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

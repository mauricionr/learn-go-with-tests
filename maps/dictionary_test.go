package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func (t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknow word", func (t *testing.T) {
		_, err := dictionary.Search("unknow")

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	dictionary.Add(word, definition)
	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("Should find added word", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

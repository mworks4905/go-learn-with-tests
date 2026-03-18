package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("Search known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	
	t.Run("Error searching for unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new word and definition", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"

		err := dictionary.Add(word, definition)
		
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Error adding an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new definition")

		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "testing"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Error updating a non-existent word", func(t *testing.T) {
		word := "testing"
		definition := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrNotFoundForUpdate)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete word", func(t *testing.T) {
		word := "testing"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("Error deleting a non-existent word", func(t *testing.T) {
		word := "testing"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)
		assertError(t, err, ErrNotFoundForDelete)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("New word should be found")
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
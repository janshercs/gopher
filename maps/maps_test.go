package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known key", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just another test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new word!")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new deets"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDefinition := "new deets"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete exiting word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "text to be deleted"}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()

	if err != want {
		t.Errorf("got error %q want %q", err, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("expected to get no errors.")
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

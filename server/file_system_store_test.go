package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader multiple times", func(t *testing.T) {
		file, removeTmpFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		defer removeTmpFile()

		store, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		got := store.GetLeague()
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("record player win and get score from a reader", func(t *testing.T) {
		file, removeTmpFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		defer removeTmpFile()

		store, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)

		want := 33
		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, want)

		store.RecordWin("Chris")

		want = 34
		got = store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		file, removeTmpFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer removeTmpFile()

		store, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEquals(t, got, want)
	})

	t.Run("works with empty file", func(t *testing.T) {
		file, removeTmpFile := createTempFile(t, "")
		defer removeTmpFile()

		_, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)
	})

	t.Run("returns league sorted", func(t *testing.T) {
		file, removeTmpFile := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer removeTmpFile()

		store, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		got := store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("did not expect an error but got one, %v", err)
	}
}

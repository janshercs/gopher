package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris") // weird to be passing pointer, but check this out for explanation https://stackoverflow.com/questions/23454940/getting-bytes-buffer-does-not-implement-io-writer-error-message/23454941

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// bytes.Buffer implements Writer interface so it can be used
	// by Greet
	buffer := bytes.Buffer{}
	Greet(&buffer, "Matt")

	got := buffer.String()
	want := "Hello, Matt"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
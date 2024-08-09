package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	var stdOut bytes.Buffer
	stdOut.ReadFrom(r)
	os.Stdout = old
	expected := "Hello, world!\n"

	if stdOut.String() != expected {
		t.Errorf("got %q, want %q", stdOut.String(), expected)
	}
}

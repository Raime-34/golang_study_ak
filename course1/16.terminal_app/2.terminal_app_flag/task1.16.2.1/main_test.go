package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFuncMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	os.Args = []string{"", "test", "0"}

	main()
	os.Stdout = old
	w.Close()

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	if stdout.String() != "├──innerDir\n|  └──otherFile\n└──someFile\n" {
		t.Errorf("printTree() = incorect output\n%s", stdout.String())
	}
}

func TestFuncMain2(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	os.Args = []string{"", "/tt", "0"}

	main()
	os.Stdout = old
	w.Close()

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	if stdout.String() != "CreateFile /tt: The system cannot find the file specified.\n" {
		t.Errorf("printTree() = incorect output\n%s", stdout.String())
	}
}

func TestFuncMain3(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	os.Args = []string{"", "/tt", "0"}

	main()
	os.Stdout = old
	w.Close()

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	if stdout.String() != "CreateFile /tt: The system cannot find the file specified.\n" {
		t.Errorf("printTree() = incorect output\n%s", stdout.String())
	}
}

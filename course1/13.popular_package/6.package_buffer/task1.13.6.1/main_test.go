package main

import (
	"bytes"
	"testing"
)

func Test_getReader(t *testing.T) {
	input := "Hi!"
	buffer := bytes.NewBufferString(input)
	b := make([]byte, 3)
	r := getReader(buffer)
	r.Read(b)
	if result := string(b); result != "Hi!" {
		t.Errorf("getReader() = %s, expected %s", result, input)
	}
}

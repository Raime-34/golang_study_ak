package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_getScanner(t *testing.T) {
	data := []byte("Hello\n,\n World!")
	buffer := bytes.NewBuffer(data)

	scanner := getScanner(buffer)

	if scanner == nil {
		t.Errorf("getSScanner() = nil expected not nil")
	}

	result := ""
	i := 0
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}

	result = strings.TrimSpace(result)

	if result != string(data) {
		fmt.Println(i)
		t.Errorf("getScanner() expected %s got %s", data, result)
	}
}

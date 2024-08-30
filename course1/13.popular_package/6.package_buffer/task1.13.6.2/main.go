package main

import (
	"bufio"
	"bytes"
)

func getScanner(buffer *bytes.Buffer) *bufio.Scanner {
	reader := bufio.NewReader(buffer)
	return bufio.NewScanner(reader)
}

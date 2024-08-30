package main

import (
	"bufio"
	"bytes"
)

func getDataString(b *bytes.Buffer) string {
	reader := bufio.NewReader(b)
	buff := make([]byte, 12)
	reader.Read(buff)
	return string(buff)
}

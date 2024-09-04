package main

import "strings"

func composePairs(pairs ...string) *strings.Reader {
	return strings.NewReader("pair=" + strings.Join(pairs, "%2C"))
}

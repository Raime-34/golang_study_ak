package main

import (
	"math/rand"
	"strings"
)

func GenerateRandomString(length int) string {
	var builder strings.Builder
	runes4Generating := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	for _ = range length {
		builder.WriteByte(runes4Generating[rand.Intn(len(runes4Generating))])
	}

	return builder.String()
}

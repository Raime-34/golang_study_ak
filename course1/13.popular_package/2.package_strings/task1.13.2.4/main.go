package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func generateActivationKey() string {
	var shards = make([]string, 4)

	for i := range shards {
		shards[i] = GenerateRandomString(4)
	}

	return strings.Join(shards, "-")
}

func GenerateRandomString(length int) string {
	var builder strings.Builder
	runes4Generating := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	for _ = range length {
		builder.WriteByte(runes4Generating[rand.Intn(len(runes4Generating))])
	}

	return builder.String()
}

func main() {
	activationKey := generateActivationKey()
	fmt.Println(activationKey)
}

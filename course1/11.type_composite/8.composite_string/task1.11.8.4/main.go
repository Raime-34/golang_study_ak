package main

import "strings"

func concatStrings(xs ...string) string {
	var builder strings.Builder

	for i := range xs {
		builder.WriteString(xs[i])
	}

	return builder.String()
}

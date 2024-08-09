package main

import (
	"testing"
	"time"
)

func Test_generateData(t *testing.T) {
	n := 30
	ch := generateData(n)

	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	counter := 0
	for range ch {
		counter++
	}

	if counter != n {
		t.Errorf("generateData(). Incorrect amount of messages in channel")
	}
}

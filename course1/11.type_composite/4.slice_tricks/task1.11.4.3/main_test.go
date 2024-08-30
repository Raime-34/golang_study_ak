package main

import (
	"testing"
)

func TestRemoveExtraMemory(t *testing.T) {
	var xs = make([]int, 0, 6)
	xs = append(xs, 1, 2, 3)

	newXs := RemoveExtraMemory(xs)

	if cap(newXs) != 3 {
		t.Errorf("the size of the allocated memory has not changed")
	}
}

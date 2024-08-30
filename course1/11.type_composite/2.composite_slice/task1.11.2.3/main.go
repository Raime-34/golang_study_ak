package main

import (
	"slices"
)

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	var dubCounters = make([]int, len(numbers))

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if bitwiseXOR(numbers[i], numbers[j]) == 0 {
				dubCounters[i]++
				dubCounters[j]++
			}
		}
	}

	resultIndex := slices.Index(dubCounters, 0)
	return numbers[resultIndex]
}

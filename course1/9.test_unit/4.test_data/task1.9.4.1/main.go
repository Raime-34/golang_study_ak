package main

import "fmt"

func main() {
	fmt.Println(average(generateSlice(10)))
}

func average(xs []float64) float64 {
	var sum float64

	for _, v := range xs {
		sum += v
	}

	return sum / float64(len(xs))
}

func generateSlice(size int) []float64 {
	result := make([]float64, size)

	for i := range size {
		result[i] = float64(i + 1)
	}

	return result
}

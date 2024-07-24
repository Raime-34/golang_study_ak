package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(*Add(5, 17))
	fmt.Println(*Max([]int{10, 5, 30, 100, -200}))
	fmt.Println(*IsPrime(5))
	fmt.Println(*IsPrime(100))
	fmt.Println(*ConcatenateStrings([]string{"строка 1 ", "строка 2 ", "... ", "строка n"}))
}

func Add(a, b int) *int {
	a += b
	return &a
}

func Max(numbers []int) *int {
	sort.Ints(numbers)
	return &numbers[len(numbers)-1]
}

func IsPrime(number int) *bool {
	result := true

	// Число простое, если делится тоько на себя и на 1
	for i := 2; i < number-1; i++ {
		if mod := number % i; mod == 0 {
			result = false
			break
		}
	}

	return &result
}

func ConcatenateStrings(strs []string) *string {
	result := []byte{}

	for i := range strs {
		result = append(result, []byte(strs[i])...)
	}

	resultStr := string(result)
	return &resultStr
}

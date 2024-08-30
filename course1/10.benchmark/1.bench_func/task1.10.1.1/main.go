package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(10))
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

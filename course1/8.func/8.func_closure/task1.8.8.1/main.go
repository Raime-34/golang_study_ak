package main

import "fmt"

var (
	count = 0
)

func main() {
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func createCounter() func() int {
	return func() int {
		count++
		return count
	}
}

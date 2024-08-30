package main

import (
	"fmt"
	"math"
)

func main() {
	a := 5
	mutate(&a)
	fmt.Println(a)

	str := "Shrek"
	ReverseString(&str)
	fmt.Println(str)
}

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {
	buff := []byte(*str)
	mid := int(math.Ceil((float64(len(*str))) / 2))

	for i := 0; i < mid; i++ {
		buff[i], buff[len(buff)-1-i] = buff[len(buff)-1-i], buff[i]
	}

	*str = string(buff)
}

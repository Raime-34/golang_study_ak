package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	isEqual, diff := CompareRoundedValues(12.3089, 12.3443453534353, 1)
	fmt.Println(isEqual)
	fmt.Println(diff)
}

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	format := fmt.Sprintf("%s.%df", "%", decimalPlaces)

	roundedAStr := fmt.Sprintf(format, a)
	roundedBStr := fmt.Sprintf(format, b)

	roundedA, err := strconv.ParseFloat(roundedAStr, 64)
	if err != nil {
		panic(err)
	}

	roundedB, err := strconv.ParseFloat(roundedBStr, 64)
	if err != nil {
		panic(err)
	}

	if roundedA == roundedB {
		isEqual = true
		difference = a
		return
	}

	difference = math.Abs(roundedA - roundedB)
	return
}

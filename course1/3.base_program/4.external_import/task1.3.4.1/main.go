package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println(DecimalSum("400.48", "101.69"))
	fmt.Println(DecimalSubtract("400.48", "101.69"))
	fmt.Println(DecimalMultiply("400.48", "101.69"))
	fmt.Println(DecimalDivide("400.48", "101.69"))
	fmt.Println(DecimalRound("400.48", 3))
	fmt.Println(DecimalGreaterThan("400.48", "101.69"))
	fmt.Println(DecimalLessThan("400.48", "101.69"))
	fmt.Println(DecimalEqual("400.48", "101.69"))
}

// Парсит строки
// если хотя бы одна из строк не содержит число вернет ошибку
func parsePair(a, b string) (A decimal.Decimal, B decimal.Decimal, err error) {
	A, err = decimal.NewFromString(a)

	if err != nil {
		return
	}

	B, err = decimal.NewFromString(b)

	return
}

func DecimalSum(a, b string) (string, error) {
	A, B, err := parsePair(a, b)

	if err != nil {
		return "", err
	}

	return A.Add(B).String(), nil
}

func DecimalSubtract(a, b string) (string, error) {
	A, B, err := parsePair(a, b)

	if err != nil {
		return "", err
	}

	return A.Sub(B).String(), nil
}

func DecimalMultiply(a, b string) (string, error) {
	A, B, err := parsePair(a, b)

	if err != nil {
		return "", err
	}

	return A.Mul(B).String(), nil
}

func DecimalDivide(a, b string) (string, error) {
	A, B, err := parsePair(a, b)

	if err != nil {
		return "", err
	}

	return A.Div(B).String(), nil
}

func DecimalRound(a string, precision int32) (string, error) {
	A, _, err := parsePair(a, "1") // Передаю 1, чтобы функция не вернула не ниловый error

	if err != nil {
		return "", err
	}

	return A.Round(precision).String(), nil
}

func DecimalGreaterThan(a, b string) (bool, error) {
	A, B, err := parsePair(a, b)

	if err != nil {
		return false, err
	}

	return A.GreaterThan(B), nil
}

func DecimalLessThan(a, b string) (bool, error) {
	A, B, err := parsePair(a, b)

	if err != nil {
		return false, err
	}

	return A.LessThan(B), nil
}

func DecimalEqual(a, b string) (bool, error) {
	A, B, err := parsePair(a, b)

	if err != nil {
		return false, err
	}

	return A.Equal(B), nil
}

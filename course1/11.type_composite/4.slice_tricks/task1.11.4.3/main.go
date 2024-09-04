package main

func RemoveExtraMemory(xs []int) []int {
	var fixedSlice = make([]int, len(xs)) // выделяю память на фактическое количество элементов в слайсе
	copy(fixedSlice, xs)
	return fixedSlice
}

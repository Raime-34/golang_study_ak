package main

func Shift(xs []int) (int, []int) {
	l := len(xs)
	var newXs = make([]int, 0, l)
	newXs = append(newXs, xs[l-1])     // добавляю в пустой слайс самый последний элемент, которы при сдвиге должен стать первым
	newXs = append(newXs, xs[:l-1]...) // добавляю остальные элементы, получаю сдвинутый слайс
	return xs[0], newXs
}

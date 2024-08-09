package main

import "strings"

func CountWordsInText(txt string, words []string) map[string]int {
	var result = make(map[string]int)
	txt = strings.ReplaceAll(strings.ToLower(txt), ".", "") // Удаление запятых
	txt = strings.ReplaceAll(strings.ToLower(txt), ",", "") // Удаление точек
	splitedTxt := strings.Fields(txt)

	for _, w := range words {
		result[w] = 0 // предварииельный перенос значений из слайса в мапу
	}

	for _, s := range splitedTxt {
		if _, ok := result[s]; ok {
			result[s]++
		}
	}

	return result
}

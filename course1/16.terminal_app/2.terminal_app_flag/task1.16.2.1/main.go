package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	var (
		path  string
		depth int64
	)

	path = args[0]
	depth, _ = strconv.ParseInt(args[1], 10, 32)

	if !strings.HasPrefix(path, "/") {
		wd, _ := os.Getwd()
		path = filepath.Join(wd, path)
	}

	printTree(path, "", true, int(depth))
}

func printTree(path string, prefix string, isLast bool, depth int) {
	info, err := os.Stat(path) // получение данных о текущей директории
	if err != nil {
		fmt.Println(err) //если ее не существует вернется ошибка
		return           // например при изначально некорректно заданной директории
	}

	// данный кейс срабатывает
	// если текущий файл или дирректория являются детьми изначальной
	// инфу о которой мы хотим узнать
	if depth > 0 {
		if isLast {
			fmt.Printf("%s└──%s\n", prefix, info.Name())
			prefix += "   "
		} else {
			fmt.Printf("%s├──%s\n", prefix, info.Name())
			prefix += "|  "
		}

	}

	// если текущий элемент является папкой
	// то необходимо рекурсивно вывести инфу о ее детях
	if info.IsDir() {
		files, _ := os.ReadDir(path)

		for i, file := range files {
			newPath := filepath.Join(path, file.Name())
			isLast := i == len(files)-1
			printTree(newPath, prefix, isLast, depth+1)
		}
	}
}

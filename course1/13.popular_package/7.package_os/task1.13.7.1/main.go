package main

import (
	"os"
	"regexp"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, perm)
	defer file.Close()

	if err != nil {
		re := regexp.MustCompile(`[^/\\]+(\.[^/\\]+)?$`) // Регулярка для поиска названия и расширения файла
		dir := re.ReplaceAllString(filePath, "")
		err = os.MkdirAll(dir, perm)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(filePath, data, perm)
	return err
}

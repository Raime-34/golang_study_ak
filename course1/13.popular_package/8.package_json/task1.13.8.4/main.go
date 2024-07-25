package main

import (
	"encoding/json"
	"os"
	"regexp"
)

func writeJSON(filePath string, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`[^/\\]+(\.[^/\\]+)?$`)
	dir := re.ReplaceAllString(filePath, "")
	err = os.MkdirAll(dir, 0644)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, bytes, 0644)
}

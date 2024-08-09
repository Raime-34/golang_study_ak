package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_getDataString(t *testing.T) {
	str := "Hello World!"
	buffer := bytes.NewBufferString(str)

	result := getDataString(buffer)

	if str != result {
		t.Errorf("getDataString() = %s, expected %s", result, str)
	}

	fmt.Println(result)
}

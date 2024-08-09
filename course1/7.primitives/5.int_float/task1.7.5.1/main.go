package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000"))
}

func binaryStringToFloat(binary string) float32 {
	n, err := strconv.ParseInt(binary, 2, 32)
	if err != nil {
		panic(err)
	}

	return *(*float32)(unsafe.Pointer(&n))
}

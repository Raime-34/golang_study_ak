package main

import (
	"fmt"
	"os/exec"
)

func ExecBin(binPath string, args ...string) string {
	out, err := exec.Command(binPath, args...).CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out)
}

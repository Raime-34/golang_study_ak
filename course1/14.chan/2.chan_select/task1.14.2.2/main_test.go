package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_timeout(t *testing.T) {
	// Вариант когда timeout отрабатывает быстрее
	timeoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)

L:
	for {
		select {
		case <-since.C:
			t.Errorf("timeout() expected true, got false")
			break L
		default:
			if timeoutFunc() {
				fmt.Println("Ok")
				break L
			}
		}
	}

	// Вариант когда быстрее time.NewTimer
	timeoutFunc = timeout(3 * time.Second)
	since = time.NewTimer(2050 * time.Millisecond)

	for {
		select {
		case <-since.C:
			fmt.Println("Ok")
			return
		default:
			if timeoutFunc() {
				t.Errorf("timeout() expected false, got true")
				return
			}
		}
	}
}

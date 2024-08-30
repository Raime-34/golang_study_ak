package main

import "time"

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	case <-time.After(1 * time.Second):
		return false
	}
}

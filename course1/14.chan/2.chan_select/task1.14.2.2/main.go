package main

import "time"

func timeout(timeout time.Duration) func() bool {
	timer := time.After(timeout)

	return func() bool {
		select {
		case <-timer:
			return true
		default:
			return false
		}
	}
}

package main

import (
	"fmt"
	"testing"
)

func Test_mergeChan(t *testing.T) {
	mainCh := make(chan int)
	var channels = make([]chan int, 5)

	for i := range channels {
		channels[i] = generateChan(i * 10)
	}

	go mergeChan(mainCh, channels...)

	counter := 0
	for n := range mainCh {
		fmt.Println(n)
		counter++
	}

	if counter != len(channels)*5 {
		t.Errorf("mergeChan(). Incorrect amount of messages in merged channel")
	}
}

func Test_mergeChan2(t *testing.T) {
	var channels = make([]chan int, 5)

	for i := range channels {
		channels[i] = generateChan(i * 10)
	}

	mainCh := mergeChan2(channels...)

	counter := 0
	for n := range mainCh {
		fmt.Println(n)
		counter++
	}

	if counter != len(channels)*5 {
		t.Errorf("mergeChan2(). Incorrect amount of messages in merged channel")
	}
}

// В канал пробрасывается 5 значений
// Первый канал в слайсе получит значения 0..4
// второй 10..14
// и т.д.
func generateChan(n int) chan int {
	ch := make(chan int)

	go func() {
		for i := n; i < n+5; i++ {
			ch <- i
		}

		close(ch)
	}()

	return ch
}

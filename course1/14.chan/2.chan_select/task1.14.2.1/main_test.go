package main

import "testing"

func Test_trySend(t *testing.T) {
	ch := make(chan int)
	defer close(ch)

	go func() {
		ch <- 5
	}()

	if trySend(ch, 50) {
		t.Errorf("trySend() = true, expected false")
	}

	<-ch // извлекаю передеанное в 10 строке значение

	go func() {
		<-ch // создаю воркера для чтения, чтобы не словить кейс с time.After
	}()

	if !trySend(ch, 50) {
		t.Errorf("trySend() = false, expected true")
	}
}

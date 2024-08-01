package main

import (
	"fmt"
	"sync"
)

func mergeChan(mergeTo chan int, from ...chan int) {
	var wg sync.WaitGroup

	for i := range from {
		wg.Add(1)
		go func(ch chan int, i int) {
			for n := range ch {
				mergeTo <- n
			}
			fmt.Printf("%v закрыт\n", i)
			wg.Done()
		}(from[i], i)
	}

	wg.Wait()
	close(mergeTo)
}

func mergeChan2(chans ...chan int) chan int {
	mergedCh := make(chan int)
	var wg sync.WaitGroup

	for i := range chans {
		wg.Add(1)

		go func(ch chan int) {
			for n := range ch {
				mergedCh <- n
			}

			wg.Done()
		}(chans[i])
	}

	// Воркер который закроет главный канал
	// во избежания дедлока
	go func() {
		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

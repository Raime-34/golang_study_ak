package main

func generateData(n int) chan int {
	ch := make(chan int)

	go func() {
		for i := range n {
			ch <- i
		}

		close(ch)
	}()

	return ch
}

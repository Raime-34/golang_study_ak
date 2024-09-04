package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime"
	"time"
)

func monitorGoroutines(prevGoroutines int) {
	time.Sleep(1 * time.Second)
	currentAmount := runtime.NumGoroutine()
	diff := float64(currentAmount) / float64(prevGoroutines)
	if diff < 0.8 {
		fmt.Println("Предупреждение: Количество горутин уменьшилось более чем на 20%")
	}
	if diff > 1.2 {
		fmt.Println("Предупреждение: Количество горутин уменьшилось более чем на 20%")
	}
	fmt.Printf("Текущее количество горутин: %v\n", currentAmount)
}

func main() {
	g, _ := errgroup.WithContext(context.Background())
	// Мониторинг горутин
	go func() {
		for {
			monitorGoroutines(runtime.NumGoroutine())
		}
	}()
	// Имитация активной работы приложения с созданием горутин
	for i := 0; i < 64; i++ {
		g.Go(func() error {
			time.Sleep(5 * time.Second)
			return nil
		})
		time.Sleep(80 * time.Millisecond)
	}
	// Ожидание завершения всех горутин
	if err := g.Wait(); err != nil {
		fmt.Println("Ошибка:", err)
	}
}

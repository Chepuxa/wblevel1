package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)

	//Завершение горутины осуществляется при получении сигнала с использованием контекста
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <- ctx.Done():
				return
			default:
				fmt.Println("Goroutine running")
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}
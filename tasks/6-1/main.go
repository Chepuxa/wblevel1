package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины. 
func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	//Завершение горутины осуществляется при получении сигнала
	go func() {
		defer wg.Done()
		for {
			select {
			case <- ch:
				return
			default:
				fmt.Println("Goroutine running")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	ch <- struct{}{}
	wg.Wait()
}
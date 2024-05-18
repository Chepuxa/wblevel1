package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		//Горутина выполняется пока канал c не закрыт
		for v := range ch {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()
}
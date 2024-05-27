package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/
func main() {
	//Задаем значение таймера (сколько будет работать программа)
	var toWait int
	fmt.Scanf("%d", &toWait)
	timer := time.After(time.Duration(toWait) * time.Second)

	var wg sync.WaitGroup

	ch := make(chan int)

	go func(in chan<- int) {
		i := 0
		//По окончанию горутины закрываем канал
		defer close(in)
		
	loop:
		for ; ; i++ {
			select {
			//Если получили сигнал от таймера, выходим из цикла
			case <-timer:
				break loop
			default:
				in <- i
			}
		}
	}(ch)

	wg.Add(1)
	go func(in <-chan int) {
		//По окончаинаю горутины сокращаем WaitGroup
		defer wg.Done()

		//Пока канал открыт, выводим данные из него
		for v := range in {
			fmt.Printf("Got %d\n", v)
		}
	}(ch)

	wg.Wait()
}

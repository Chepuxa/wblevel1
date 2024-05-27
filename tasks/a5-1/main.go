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
	var timer int
	fmt.Scanf("%d", &timer)

	var wg sync.WaitGroup

	ch := make(chan int)

	go func(in chan<- int) {
		//По окончанию горутины закрываем канал
		defer close(in)

		//Пока таймер !=0, уменьшаем его каждую секунду
		//Каждый цикл инкрементируем i и отправляем в канал
		for i := 1; timer != 0; timer, i = timer-1, i+1 {
			in <- i
			time.Sleep(time.Second)
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

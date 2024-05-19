package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
func main() {
	var workers int
	fmt.Scanf("%d", &workers)

	c := make(chan int)

	var stopChan = make(chan os.Signal, 2)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		//По окончанию горутины закрываем канал
		defer close(c)
	
	L:
		for {
			select {
			//Если получили сигнал os.Interrupt, завершить цикл for
			case <-stopChan:
				break L
			default:
				c <- rand.Int()
			}
		}
	}()

	for i := 1; i <= workers; i++ {
		go func(i int) {
			//По окончаинаю горутины сокращаем WaitGroup
			defer wg.Done()

			//Считываем данные из канала c до тех пор, пока он не закрылся
			for v := range c {
				fmt.Printf("%d got %d\n", i, v)
			}
		}(i)
	}
	signal.Notify(stopChan, os.Interrupt)
	wg.Wait()
}

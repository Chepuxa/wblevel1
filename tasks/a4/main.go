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

	intChan := make(chan int, workers)

	var stopChan = make(chan os.Signal, 2)

	var wg sync.WaitGroup

	go func(intCh chan<- int, stopCh <-chan os.Signal) {
		//По окончанию горутины закрываем канал
		defer close(intCh)
	
	L:
		for {
			select {
			//Если получили сигнал os.Interrupt, завершить цикл for
			case <-stopCh:
				break L
			default:
				intCh <- rand.Int()
			}
		}
	}(intChan, stopChan)

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func(i int, intCh <-chan int) {
			//По окончаинаю горутины сокращаем WaitGroup
			defer wg.Done()

			//Считываем данные из канала c до тех пор, пока он не закрылся
			for v := range intCh {
				fmt.Printf("\n%d got %d\n", i, v)
			}
		}(i, intChan)
	}
	/*
	Для остановки всех воркеров используется сигнал os.Interrupt
	При получении такого сигнала, отправка данных в канал intCha прекратится и канал закроется
	При закрытии канала, считывающие горутины закончат свои работы, вызвав отложенный wg.Done()
	Когда все горутины завершатся, wg опустеет и программа завершится
	Такой подход гарантирует, что все данные из основного потока будут считаны в случае завершения программы
	*/
	signal.Notify(stopChan, os.Interrupt)
	wg.Wait()
}

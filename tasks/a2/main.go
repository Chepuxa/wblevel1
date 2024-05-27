package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
func main() {
	var wg sync.WaitGroup
	in := []int{2, 4, 6, 8, 10}
	out := make(chan int)

	for _, v := range in {
		wg.Add(1)

		go func (j int) {
			fmt.Println(j * j)
			wg.Done()
		}(v)
	}

	wg.Wait()
	close(out)
}
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
func main() {
    var ai atomic.Int32
    var wg sync.WaitGroup
	in := []int32{2, 4, 6, 8, 10}
    
    for _, v := range in {
        wg.Add(1)

        go func(i int32) {
            ai.Add(i * i)
            wg.Done()
        }(v)
    }

    wg.Wait()

	fmt.Println(ai.Load())
}
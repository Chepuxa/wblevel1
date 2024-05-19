package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
func main() {
	workers := 10
	increment := 100
	result := 10 * 100

    var at atomic.Uint64
    var wg sync.WaitGroup

    for i := 0; i < workers; i++ {
        wg.Add(1)

        go func() {
			defer wg.Done()

            for j := 0; j < increment; j++ {

                at.Add(1)
            }
        }()
    }

    wg.Wait()

    fmt.Println(at.Load())
	fmt.Println(at.Load() == uint64(result))
}
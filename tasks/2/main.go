package main

import "fmt"

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
func main() {
	in := []int{2, 4, 6, 8, 10}
	out := make(chan int)
    go func() {
		//Передаем в канал возведенные в квадрат числа из массива in
        for _, n := range in {
            out <- n * n
        }
        close(out)
    }()
    
	for v := range out {
		fmt.Println(v)
	}
}
package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. 
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/
func main() {
	a := make(chan int)
	b := make (chan int)
	ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() {
		defer close(a)
		for _, v := range ar {
			a <- v
		}
	}()

	go func() {
		defer close(b)
		for v := range a {
			b <- v * 2
		}
	}()

	for v := range b {
		fmt.Println(v)
	}
}
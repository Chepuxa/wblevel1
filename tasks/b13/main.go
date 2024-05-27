package main

import "fmt"

//Поменять местами два числа без создания временной переменной.
func main() {
	i, v := 10, 20
	i, v = v, i
	fmt.Printf("%d, %d", i, v)
}
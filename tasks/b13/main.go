package main

import "fmt"

//Поменять местами два числа без создания временной переменной.
func main() {
	i := 10
	v := 20
	i, v = v, i
	fmt.Printf("%d, %d", i, v)
}
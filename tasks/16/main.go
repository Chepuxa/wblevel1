package main

import (
	"fmt"
	"math/rand"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
func main() {
	i := []int{12, 2, 44, 10, 1, 13, 8, 9, 100, 36, 4, 23, 18, 5, 39, 6}
	fmt.Println(quickSort(i))
}

func quickSort(n []int) []int {
	if len(n) < 2 {
		return n
	}

	l := []int{}
	r := []int{}
	j := n[rand.Intn(len(n))]
	for _, v := range n {
		if v < j {
			l = append(l, v)
		} else if v > j {
			r = append(r, v)
		}
	}
	l = append(l, j)

	return append(quickSort(l), quickSort(r)...)
}
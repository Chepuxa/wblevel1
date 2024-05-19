package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := []int{1, 2, 3, 4, 5, 6}
	n := rand.Intn(len(i))
	fmt.Printf("Remove index: %d, element at index: %d\n", n, i[n])
	rm := removeIndex(i, n)
	fmt.Printf("Original slice: %v\n", i)
	fmt.Printf("Slice with removed element: %v\n", rm)
}

func removeIndex(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}
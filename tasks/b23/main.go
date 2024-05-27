package main

import (
	"fmt"
	"math/rand"
)

//Удалить i-ый элемент из слайса.
func main() {
	i := []int{1, 2, 3, 4, 5, 6}
	n := rand.Intn(len(i))
	fmt.Printf("Remove index: %d, element at index: %d\n", n, i[n])
	fmt.Printf("Original slice: %v\n", i)
	i = removeIndex(i, n)
	fmt.Printf("Slice after removal\n: %v\n", i)
	fmt.Printf("Cap: %d, Len: %d\n", cap(i), len(i))
}

func removeIndex(s []int, i int) []int {
    return append(s[:i], s[i+1:]...)
}
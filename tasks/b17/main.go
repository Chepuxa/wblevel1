package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.
func main() {
	n := []int{1, 2, 4, 5, 6, 8, 9, 10, 12, 13, 18, 23, 36, 39, 44, 100}
	fmt.Println(binarySearch(n, 100))
}

func binarySearch(n []int, i int) (int, bool) {
	low := 0
	high := len(n)

	for low <= high {
		mid := (low + high) / 2
		if guess := n[mid]; guess > i {
			high = mid - 1
		} else if guess < i {
			low = mid + 1
		} else {
			return mid, true
		}
	}
	return 0, false
}
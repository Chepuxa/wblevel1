package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.
func main() {
	a := []int{4, 7, 3, 12, 6, 5, 10, 9}
	b := []int{3, 20, 45, 4, 100, 31, 6}
	res := []int{}
	m := make(map[int]struct{})

	for _, v := range a {
		m[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}

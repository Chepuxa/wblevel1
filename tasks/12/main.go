package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func main() {
	i := []string{"cat", "cat", "dog", "cat", "tree"}
	res := []string{}
	m := make(map[string]bool)
	for _, v := range i {
		if _, ok := m[v] ; !ok {
			res = append(res, v)
		}
		m[v] = true
	}

	fmt.Println(res)
}
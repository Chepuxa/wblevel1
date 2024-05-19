package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
func main() {
	s := "a😏sя😒ы"
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	a := []rune(s)
	for l, r := 0, len(a) - 1; l < r; l, r = l + 1, r - 1 {
		a[l], a[r] = a[r], a[l]
	}
	return string(a)
}
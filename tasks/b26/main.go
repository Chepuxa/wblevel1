package main

import (
	"fmt"
	"unicode"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false
*/
func main() {
	s := "_-!abcA"
	fmt.Println(unique(s))
}

func unique(s string) bool {
	m := make(map[rune]struct{})
	for _, v := range s {
		v = unicode.ToLower(v)
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
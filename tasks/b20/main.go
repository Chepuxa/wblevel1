package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Разработать программу, которая переворачивает слова в строке. 
Пример: «snow dog sun — sun dog snow».
*/
func main() {
	s := "машина  😏sя😒   sun кот flow                          axiom ведро nasa melon при-вет"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	n := []string{}
	j := []rune{}
	var prev rune
	for _, v := range s {
		//Если текущий символ - space и предыдущей !space - добавляем собранное слово из j в набор строк n и чистим j
		if unicode.IsSpace(v) && !unicode.IsSpace(prev) {
			n = append(n, string(j))
			clear(j)
		//Если текущий символ !space - добавляем сивол в j
		} else if !unicode.IsSpace(v) {
			j = append(j, v)
		}
		prev = v
	}
	n = append(n, string(j))

	//Переворачиваем набор строк n
	for l, r := 0, len(n) - 1; l < r; l, r = l + 1, r - 1 {
		n[l], n[r] = n[r], n[l]
	}
	//Объединяем набор строк n в строку с разделителем space
	return strings.Join(n, " ")
}
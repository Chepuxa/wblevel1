package main

import "fmt"

/*
Вопрос: Что выведет данная программа и почему?

	func update(p *int) {
	  b := 2
	  p = &b
	}

	func main() {
	  var (
	     a = 1
	     p = &a
	  )
	  fmt.Println(*p)
	  update(p)
	  fmt.Println(*p)
	}

Ответ: 1
Функция оперирует с копиями переданных аргументов.
В функции update, параметр p, над которым производится операция, является копией аргумента p из main().
При изменении ссылки, которую хранит p в методе update(), ссылка, которую хранит p в методе main() останется той же
Значение p из main() можно изменить путем "дереференсинга" - изменив значение, на которое ссылается p, не меняя саму ссылку
*/
func update(p *int) {
	b := 2
	*p = b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	fmt.Println(p)
	update(p)
	fmt.Println(*p)
	fmt.Println(p)
}

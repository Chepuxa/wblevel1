package main

import "fmt"

/*
Вопрос: Что выведет данная программа и почему?


func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}

Ответ:
[b, b, a]
[a, a]
Операция append в анонимной функции будет вынуждена нарастить cap слайса, выделив под это память для нового массива, и с этого момента переменная slice будет указывать на другой массив
Все последующие операции изменяют новосозданный массив, не затрагивая внешний slice
*/
func main() {
	slice := []string{"a", "a"}
  
	func(slice []string) {
	   slice = append(slice, "a")
	   slice[0] = "b"
	   slice[1] = "b"
	   fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
  }
  
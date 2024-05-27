package main

import (
	"fmt"
	"math/rand"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Ответ: Срез по строке не аллоцирует новую строку, и новый срез [байт, лежащий в основе строке] будет указывать на массив байт созданной огромной строки
Т.к. justString глобальная переменная, v по итогу "убежит" в хип и будет там висеть, хотя нам нужны только 0-99 элементы оттуда
Вторая проблема возникнет, если в строке v будут содержаться символы, занимающие больше 1 байта, в следствии чего срез может "распилить" символ и justString будет содержать некорректные данные
*/

var justString string

func main() {
	n := 100
	someFunc(n)
	fmt.Println(len([]rune(justString)) == n)
}

func someFunc(n int) {
	v := createHugeString(1 << 10)
	justString = string(copyBytes([]rune(v), n))
}

func createHugeString(n int) string {
	const letters = "qwertnmйц😏😒🙂‍↔️😞😔укенгшмитьб😁ю世🖖🖖界"
	runes := []rune(letters)
	buffer := make([]rune, n)
	for b := range buffer {
		buffer[b] = runes[rand.Intn(len(runes))]
	}
	return string(buffer)
}

func copyBytes(s []rune, n int) []rune {
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = s[i]
	}
	return res
}
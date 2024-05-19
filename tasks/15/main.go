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

Ответ: т.к. мы глобальной переменной justString задаем значение, равное срезу по локальной переменной v, внутри этого среза будет храниться ссылка на v
Из-за этого v будет храниться в памяти после выполнения функции someFunc
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
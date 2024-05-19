package main

import (
	"fmt"
	"unsafe"
)

/*
Вопрос: Какой размер у структуры struct{}{}?

Ответ: 
0, т.к. это адресс в памяти который не указывает ни на какие данные
Размер struct = размер состовляющих типов + отступы
Более того, любой указатель на struct{}{} или тип struct{} всегда указывает на один и тот же адресс
Также struct, состоящий из пустых struct, будет занимать 0 байт
 */

type A struct {
	A byte
	B byte
}

type B struct {
	A byte
}

type C struct{}

type D struct {
	Da struct{}
	Db struct{}
}

func main() {
	fmt.Println(unsafe.Sizeof(A{}))
	fmt.Println(unsafe.Sizeof(B{}))
	fmt.Println(unsafe.Sizeof(C{}))
	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(unsafe.Sizeof(D{}))

	fmt.Printf("%p\n", &C{})
	fmt.Printf("%p\n", &struct{}{})
	fmt.Printf("%p\n", &D{})

}
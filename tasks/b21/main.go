package main

import "fmt"

//Внешние структуры и методы, которые мы не можем изменять
type A struct {}

func (a *A) DoA() {
    fmt.Println("Doing A")
}

type B struct {}

func (b *B) DoB(bs string) {
    fmt.Printf("Doing %s\n", bs)
}

//Интерфейс адаптера
type Adapter interface {
    Do()
}

//Реализация адаптера для A
type AdapterA struct{
    *A
}

func (adapter *AdapterA) Do() {
    adapter.DoA()
}

func NewAdapterA(a *A) Adapter {
    return &AdapterA{a}
}

//Реализация адаптера для B
type AdapterB struct{
    *B
}

func (adapter *AdapterB) Do() {
    adapter.DoB("B")
}

func NewAdapterB(b *B) Adapter {
    return &AdapterB{b}
}

func main() {
	a := NewAdapterA(&A{})
	b := NewAdapterB(&B{})
	a.Do()
	b.Do()
}
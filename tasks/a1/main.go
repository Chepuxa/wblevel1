package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
type Human struct {
	Hands int
	Legs  int
	Head  int
	WalkSpeed int
}

//Метод структуры Human
func (h *Human) Walk(time int) {
	fmt.Printf("Covered distance: %d\n", time * h.WalkSpeed)
}

//Композиция
type Action struct {
	Human
}

func main() {

	a := Action{
		Human{
			Hands: 2,
			Legs:  2,
			Head:  1,
			WalkSpeed: 3,
		}}

	fmt.Printf("Hands amount: %d\n", a.Hands)
	a.Walk(5)
}

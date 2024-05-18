package main

import "fmt"

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

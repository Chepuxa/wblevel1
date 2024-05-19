package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
func main() {
	ch := make(chan struct{})
	a := []interface{}{10, "asd", true, ch}

	for _, v := range a {
		switch v.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		case chan struct{}:
			fmt.Println("chan")
		}
	}
}
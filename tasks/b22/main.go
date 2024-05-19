package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("1000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
	b.SetString("1000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
	fmt.Printf("%v, %v\n", a, b)

	a.Mul(a, b)
	fmt.Println(a)
	a.Div(a, b)
	fmt.Println(a)
	a.Add(a, b)
	fmt.Println(a)
	a.Sub(a, b)
	fmt.Println(a)
}
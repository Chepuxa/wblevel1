package main

import "fmt"

func main() {
	// var num int64
	// var i int64
	// fmt.Println("Enter decimal number")
	// fmt.Scanf("%d", &num)
	// fmt.Printf("Binary representation: %b\n", num)
	// fmt.Println("Enter position to change to 1")
	// fmt.Scanf("%d", &i)

	// num |= (1 << i)
    // fmt.Println(num)

	v := 2
	fmt.Printf("%b\n", v)
	v |= v
	fmt.Printf("%b\n", v)
	// i := v << 2
	// fmt.Printf("%b\n", i)
}
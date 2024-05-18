package main

import "fmt"

func main() {
	in := []int{2, 4, 6, 8, 10}
	out := make(chan int)
    go func() {
		sum := 0
        for _, n := range in {
            sum += n * n
        }
		out <- sum
        close(out)
    }()
    
	fmt.Println(<-out)
}
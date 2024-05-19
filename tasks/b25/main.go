package main

import (
	"fmt"
	"time"
)

func main() {
	Sleep(5)
	fmt.Println("Done")
}

func Sleep(x int) {
    <-time.After(time.Second * time.Duration(x))
}
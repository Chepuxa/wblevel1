package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func main() {
	pa := NewPoint(0, 1)
	pb := NewPoint(2, -2)
	fmt.Println(findDistance(pa, pb))
}

func findDistance(a *Point, b *Point) float64 {
	return math.Sqrt(math.Pow(b.x - a.x, 2) + math.Pow(b.y - a.y, 2))
}
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
	return &Point{
		x, y,
	}
}

func Distance(first *Point, second *Point) float64 {
	distanceX := math.Abs(first.x - second.x)
	distanceY := math.Abs(first.y - second.y)
	return math.Sqrt(math.Pow(distanceX, 2) + math.Pow(distanceY, 2))
}

func main() {
	firstPoint := NewPoint(-10, 5.6)
	secondPoint := NewPoint(3.4, 1.5)
	distance := Distance(firstPoint, secondPoint)
	fmt.Printf("distance between pointers %f", distance)
}

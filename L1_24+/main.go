package main

import (
	"fmt"
	"math"
)

func main() {
	pointOne, pointTwo := NewPoint(5, -4), NewPoint(56, 71)
	distance := DistanceBetweenPoints(pointOne, pointTwo)
	fmt.Printf("Расстояние между точками = %f ", distance)
}

type point struct {
	x int
	y int
}

func NewPoint(x int, y int) *point {
	return &point{
		x: x,
		y: y,
	}
}

// находим по формуле корень из квадрата разности координат
func DistanceBetweenPoints(pointOne, pointTwo *point) float64 {
	dx := pointTwo.x - pointOne.x
	dy := pointTwo.y - pointOne.y
	distance := math.Sqrt(float64(dx*dx + dy*dy))
	return distance
}

// 24. Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

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

func computeDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(2, 2)

	fmt.Println("First point:", *p1)
	fmt.Println("Second point", *p2)
	fmt.Println()
	fmt.Println("Distance:", computeDistance(p1, p2))

}

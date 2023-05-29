package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 Point) float64 {
	// длина первого катета
	w := math.Abs(p1.x - p2.x)
	// длина второго катета
	h := math.Abs(p1.y - p2.y)

	return math.Sqrt(math.Pow(w, 2) + math.Pow(h, 2))
}


func main() {
	p := NewPoint(1.213, 2.2)
	p1 := NewPoint(1.23, 2.19)

	// расстояние между точками
	fmt.Println(Distance(*p, *p1))
}

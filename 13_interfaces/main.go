package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rect struct {
	width, height float64
}

type Triangle struct {
	base, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (t Triangle) area() float64 {
	return (t.base * t.height) / 2
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rect{width: 10, height: 5}
	triangle := Triangle{base: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
	fmt.Printf("Rectangle Area: %f\n", getArea(triangle))

}

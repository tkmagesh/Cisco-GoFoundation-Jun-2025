package main

import (
	"fmt"
	"math"
)

// ver 1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ver 2.0
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// ver 3.0
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// ver 4.0
// Print the "perimeter" for all the above shapes
// Circle : 2 * pi * r
// Rectangle : 2 * (h + w)
// Square : 4 * side

// Make sure NONE of the above code is modified
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

type PerimeterFinder interface {
	Perimeter() float64
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// ver 5.0
/*
func PrintShape(x interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}) {
	PrintArea(x)      // x should be interface{Area() float64}
	PrintPerimeter(x) // x should be interface{ Perimeter() float64}
}
*/

/*
func PrintShape(x interface {
	AreaFinder
	PerimeterFinder
}) {
	PrintArea(x)      // x should be interface{Area() float64}
	PrintPerimeter(x) // x should be interface{ Perimeter() float64}
}
*/

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x ShapeStatsFinder) {
	PrintArea(x)      // x should be interface{Area() float64}
	PrintPerimeter(x) // x should be interface{ Perimeter() float64}
}

func main() {
	c := Circle{Radius: 16}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 16, Width: 15}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)

	s := Square{Side: 22}
	/*
		PrintArea(s)
		PrintPerimeter(s)
	*/
	PrintShape(s)

}

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

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("error : parameter does not support Area() method")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("error : parameter does not support Area() method")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
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

func main() {
	c := Circle{Radius: 16}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)

	r := Rectangle{Height: 16, Width: 15}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)

	s := Square{Side: 22}
	PrintArea(s)

	PrintArea(100)
}

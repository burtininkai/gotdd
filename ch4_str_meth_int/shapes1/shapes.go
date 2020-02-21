package main

import "math"

//Shape interface
type Shape interface {
	Area() float64
}

//Rectangle contains info about rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

//Area returns are asize of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle contains info about circle
type Circle struct {
	Radius float64
}

//Area returns are asize of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns perimeter value of rectangle
func Perimeter(rectangle Rectangle) float64 {
	var perimeter = 2 * (rectangle.Width + rectangle.Height)
	return perimeter
}

//Area returns are size of rectangle
func Area(rectangle Rectangle) float64 {
	var area = rectangle.Width * rectangle.Height
	return area
}

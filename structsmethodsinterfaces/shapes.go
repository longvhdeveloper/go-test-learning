package structsmethodsinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Edge   float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (triangle Triangle) Area() float64 {
	return (triangle.Edge * triangle.Height) / 2
}

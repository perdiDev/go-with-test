package structmethodinterface

import "math"

func Perimeter(shape Rectangle) float64 {
	return 2 * (shape.Width + shape.Height)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Width float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Width * t.Height)/2
}

type Shape interface {
	Area() float64
}

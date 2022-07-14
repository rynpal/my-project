package main

import "math"

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) (area float64) {
	return rectangle.Height * rectangle.Width
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

type Triangle struct{
	Base float64
	Height float64
}

func (t Triangle) Area() float64  {
	return (t.Base * t.Height) / 2
}

type Shape interface {
	Area() float64
}

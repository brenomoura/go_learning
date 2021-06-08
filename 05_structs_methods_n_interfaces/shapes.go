package main

import "math"

type Rectangle struct {
	Weight float64
	Height float64
}

type Triangle struct {
	Weight float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Weight)
}

func Area(height, weight float64) float64 {
	return height * weight
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Weight
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Weight) / 2
}

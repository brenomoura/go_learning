package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangule := Rectangle{10.0, 10.0}
	got := Perimeter(rectangule)
	want := 40.0

	if got != want {
		t.Errorf("Got: %.2f, Want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v Got: %.2f, Want: %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}

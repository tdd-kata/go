package _struct // 패키지명이 키워드와 동일할 때

import "math"

// Shape is implemented by anything that can tell us its Area.
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// perimeter returns the perimeter of a rectangle.
func perimeter(rectangle Rectangle) float64 {
	// symbol이 소문자로 시작한다면 정의된 패키지 밖에서는 private
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents the dimensions of a triangle.
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the triangle.
func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

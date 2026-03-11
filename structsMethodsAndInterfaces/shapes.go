package shapes

import "math"

type Rectangle struct {
	Height float64
	Width float64
}

func (r Rectangle) Area () float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area () float64 {
	return math.Pi * (c.Radius * c.Radius)
}

type Triangle struct {
	Height float64
	Base float64
}

func (t Triangle) Area () float64 {
	return t.Base * t.Height / 2
}


func Perimeter(height, width float64) float64 {
	return 2 * (height + width)
}

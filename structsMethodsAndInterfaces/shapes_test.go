package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

type Shape interface{
	Area() float64
}

func TestArea(t *testing.T) {

	areaTest := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 12.0, Width: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 5.0}, hasArea: 78.53981633974483},
		{name: "Triangle", shape: Triangle{Height: 12, Base: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v - got area: %g, want area: %g", tt.shape, got, tt.hasArea)
			}
		})
	}
	

	// Old Way
	// ============
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g, want %g", got, want)
	// 	}
	// }
	// t.Run("Rectangle Area", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })
	// t.Run("Cirle Area", func (t *testing.T)  {
	// 	circle := Circle{5.0}
	// 	checkArea(t, circle, 78.53981633974483)
	// })
}
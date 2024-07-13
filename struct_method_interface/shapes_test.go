package structmethodinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// checkArea:= func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got:= shape.Area()

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// }
	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle:=Circle{10.0}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

	areaTest := []struct {
		name string
		shape Shape
		hasArea float64
	} {
		{name: "rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{3.0, 4.0}, hasArea: 6.0},
	}

	for _, tt:=range areaTest {
		got:=tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g",tt.shape, got, tt.hasArea)
		}
	}
}
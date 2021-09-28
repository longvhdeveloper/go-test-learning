package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f want %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				Width: 15.0, Height: 10.0},
			expected: 150.0,
		},
		{
			name:     "Circle",
			shape:    Circle{Radius: 10.0},
			expected: 314.1592653589793,
		},
		{
			name:     "Triangle",
			shape:    Triangle{12, 6},
			expected: 36.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expected {
				t.Errorf("got %g want %g", got, tt.expected)
			}
		})
	}

	//checkAreas := func(t testing.TB, shape Shape, expected float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != expected {
	//		t.Errorf("got %g want %g", got, expected)
	//	}
	//}

	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{
	//		Width:  15.0,
	//		Height: 10.0,
	//	}
	//	expected := 150.0
	//	checkAreas(t, rectangle, expected)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{Radius: 10.0}
	//	expected := 314.1592653589793
	//	checkAreas(t, circle, expected)
	//})
}

package main

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
	/*
		Table driven test is a great tool
		when you need to test various implementations of an interface
	*/

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 8.0}, 80.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2g want %.2g", got, tt.want)
		}
	}
}

/*
In Go interface resolution is implicit.
If the type you pass in matches what the interface is asking for,
it will compile
*/

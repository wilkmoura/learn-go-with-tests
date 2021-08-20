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
	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 8.0}
		got := rectangle.Area()
		want := 80.0

		if got != want {
			t.Errorf("got %.2g want %.2g", got, want)
		}
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2g want %.2g", got, want)
		}
	})
}

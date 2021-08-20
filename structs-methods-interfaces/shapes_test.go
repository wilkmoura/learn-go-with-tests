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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2g want %.2g", got, want)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 8.0}
		want := 80.0
		checkArea(t, rectangle, want)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}

/*
In Go interface resolution is implicit.
If the type you pass in matches what the interface is asking for,
it will compile
*/

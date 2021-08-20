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
	rectangle := Rectangle{10.0, 8.0}
	got := Area(rectangle)
	want := 80.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

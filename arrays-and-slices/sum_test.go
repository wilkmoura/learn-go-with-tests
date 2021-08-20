package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

	t.Run("Collection of varying size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

/*
It is important to question the value of your tests.
It should not be a goal to have as many tests as possible,
but rather to have as much confidence as possible in your code base.
Having too many tests can turn in to a real problem and it just adds more overhead in maintenance.
EVERY TEST HAS A COST.
*/

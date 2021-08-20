package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of varying size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5})
	want := []int{6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

/*
It is important to question the value of your tests.
It should not be a goal to have as many tests as possible,
but rather to have as much confidence as possible in your code base.
Having too many tests can turn in to a real problem and it just adds more overhead in maintenance.
EVERY TEST HAS A COST.
*/

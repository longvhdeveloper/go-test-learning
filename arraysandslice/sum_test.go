package arraysandslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		expected := 10

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})
}

func checkSums(t testing.TB, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	checkSums(t, got, expected)
}

func TestSumTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})

	t.Run("safety sum with empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{0, 9})
		expected := []int{0, 9}

		checkSums(t, got, expected)
	})
}

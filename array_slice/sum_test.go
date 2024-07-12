package arrayslice

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("Got %d expected %d given %v with type %T", sum, expected, numbers, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	sumAll := SumAll([]int{1,2,3}, []int{1, 2, 3, 4})
	expected:= []int{6, 10}

	if !slices.Equal(sumAll, expected) {
		t.Errorf("Got %v expected %v", sumAll, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func (t testing.TB, want, got []int)  {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}		
	}
	t.Run("make the sum of slice", func(t *testing.T) {
		got := SumAllTails([]int{2, 3}, []int{0, 9})
		want := []int{3, 9}

		checkSum(t, got, want)
	})
	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSum(t, got, want)
	})
}
package arraysandslices

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		nums := []int{1,2,3}
		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, nums)
		}
	})

	t.Run("Sum all collections", func(t *testing.T) {
		// numSlices := [][]int{{1,2}, {4,5}}
		got := SumAll([]int{1,2}, []int{4,5})
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTeils(t *testing.T) {

	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Sum all collection tails", func(t *testing.T) {
		got := SumAllTails([]int{1,2,3}, []int{4,5,6})
		want :=[]int{5, 11}

		checkSum(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4,5,6,7})
		want :=[]int{0, 18}

		checkSum(t, got, want)
	})
}

func TestMasterSummer(t *testing.T) {
	got := MasterSum([]int{1,2,3}, []int{4,5,6})
	want := 21
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
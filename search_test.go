package search

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("search 10 in [1, 3, 5, 10, 12, 15]", func(t *testing.T) {
		arr := []int{1, 3, 5, 10, 12, 15}
		x := 10

		got, _ := BinarySearch(arr, x)

		want := 3
		if got != want {
			t.Errorf("BinarySearch failed, expected %d got %d", want, got)
		}
	})
	t.Run("search 15 in [1, 3, 5, 10, 12, 15]", func(t *testing.T) {
		arr := []int{1, 3, 5, 10, 12, 15}
		x := 15

		got, _ := BinarySearch(arr, x)

		want := 5
		if got != want {
			t.Errorf("BinarySearch failed, expected %d got %d", want, got)
		}
	})
	t.Run("search 30 in [1, 3, 5, 10, 12, 15]", func(t *testing.T) {
		arr := []int{1, 3, 5, 10, 12, 15}
		x := 30

		_, got := BinarySearch(arr, x)

		want := false
		if got != want {
			t.Errorf("BinarySearch failed, expected %t got %t", want, got)
		}
	})
}

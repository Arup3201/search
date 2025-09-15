package search

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("search 10 in [1, 3, 5, 10, 12, 15]", func(t *testing.T) {
		arr := []int{1, 3, 5, 10, 12, 15}
		x := 10

		got := BinarySearch(arr, x)

		want := 3
		if got != want {
			t.Errorf("BinarySearch failed, expected %d got %d", want, got)
		}
	})
}

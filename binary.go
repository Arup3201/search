package search

func BinarySearch(nums []int, target int) (int, bool) {
	n := len(nums)
	l, h := 0, n-1

	for l <= h {
		mid := l + (h-l)/2

		if nums[mid] == target {
			return mid, true
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}

	return -1, false
}

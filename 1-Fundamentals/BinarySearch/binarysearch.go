package BinarySearch

//BinarySearch returns the smallest index of a number appearing in a sorted list
func BinarySearch(object []int, target int) int {
	if len(object) == 0 {
		return -1
	}
	left := 0
	right := len(object) - 1
	for left <= right {
		median := left + (right-left)/2
		if target < object[median] {
			right = median - 1
		} else if target > object[median] {
			left = median + 1
		} else {
			return median
		}
	}
	return -1
}

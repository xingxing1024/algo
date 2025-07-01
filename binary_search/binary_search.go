package binarysearch

import "sort"

func BinarySearchFirst(data []int, foundVal int) int {
	sort.Ints(data)

	l := 0
	r := len(data) - 1
	result := -1
	for l <= r {
		mid := int((l + r) / 2)
		if data[mid] == foundVal {
			result = mid
			r = mid - 1
		} else if foundVal > data[mid] {
			l = mid + 1
		} else if foundVal < data[mid] {
			r = mid - 1
		}
	}
	return result
}

func BinarySearchLast(data []int, foundVal int) int {
	sort.Ints(data)

	l := 0
	r := len(data) - 1
	result := -1
	for l <= r {
		mid := int((l + r) / 2)
		if data[mid] == foundVal {
			result = mid
			l = mid + 1
		} else if foundVal > data[mid] {
			l = mid + 1
		} else if foundVal < data[mid] {
			r = mid - 1
		}
	}
	return result
}

package search

import (
	"math"

	ds "github.com/DavidCai1993/datastructures.go"
)

// BinarySearch finds the given element's location by binary search algorithm.
// The list given must be sorted.
// Time complexity O(logn)
// Space complexity O(1)
func BinarySearch(cs ds.Comparables, val ds.Comparable) int {
	if len(cs) == 0 {
		return -1
	}

	lower, upper := 0, len(cs)-1

	for upper-lower > 1 {
		mid := lower + int(math.Floor(float64(upper-lower)/2))

		c := val.Compare(cs[mid])

		if c == 0 {
			return mid
		} else if c < 0 {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}

	// upper - lower == 1
	if cs[lower].Compare(val) == 0 {
		return lower
	}

	if cs[upper].Compare(val) == 0 {
		return upper
	}

	return -1
}

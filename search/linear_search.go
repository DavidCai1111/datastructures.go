package search

import ds "github.com/DavidCai1993/datastructures.go"

// LinearSearch finds the given element's location by linear search algorithm.
// Time complexity O(n)
// Space complexity O(1)
func LinearSearch(cs ds.Comparables, val ds.Comparable) int {
	for i, c := range cs {
		if c.Compare(val) == 0 {
			return i
		}
	}

	return -1
}

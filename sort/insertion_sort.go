package sort

import (
	ds "github.com/DavidCai1993/datastructures.go"
)

// InsertionSort sorts the given []Comparable by insertion sort algorithm.
// Time complexity O(n^2)
// Space complexity O(1)
func InsertionSort(c ds.Comparables) ds.Comparables {
	for i := 0; i < len(c)-1; i++ {
		k := i + 1
		for j := i; j >= 0; j-- {
			if c[k].Compare(c[j]) < 0 {
				c[k], c[j] = c[j], c[k]
				k = j
			}
		}
	}

	return c
}

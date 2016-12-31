package sort

import (
	ds "github.com/DavidCai1993/datastructures.go"
)

// BubbleSort sorts the given []Comparable by bubble sort algorithm.
// Time complexity O(n^2)
// Space complexity O(1)
func BubbleSort(c ds.Comparables) ds.Comparables {
	for i := 0; i < len(c); i++ {
		for j := 1; j < len(c); j++ {
			if c[j].Compare(c[j-1]) < 0 {
				c[j], c[j-1] = c[j-1], c[j]
			}
		}
	}

	return c
}

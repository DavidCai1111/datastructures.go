package sort

import ds "github.com/DavidCai1993/datastructures.go"

// SelectionSort sorts the given []Comparable by selection sort algorithm.
// Time complexity O(n^2)
// Space complexity O(1)
func SelectionSort(c ds.Comparables) ds.Comparables {
	for i := 0; i < len(c)-1; i++ {
		min := i
		for j := i + 1; j < len(c); j++ {
			if c[j].Compare(c[min]) < 0 {
				min = j
			}
		}

		c[i], c[min] = c[min], c[i]
	}

	return c
}

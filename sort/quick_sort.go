package sort

import ds "github.com/DavidCai1993/datastructures.go"

// QuickSort sorts the given []Comparable by quick sort algorithm.
// Time complexity O(nlogn)
// Space complexity O(n)
func QuickSort(c ds.Comparables) ds.Comparables {
	if len(c) <= 1 {
		return c
	}

	left, right := make(ds.Comparables, 0), make(ds.Comparables, 0)

	for i := 1; i < len(c); i++ {
		if c[i].Compare(c[0]) < 0 {
			left = append(left, c[i])
		} else {
			right = append(right, c[i])
		}
	}

	return append(append(QuickSort(left), c[0]), QuickSort(right)...)
}

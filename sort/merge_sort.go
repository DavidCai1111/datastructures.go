package sort

import (
	"math"

	ds "github.com/DavidCai1993/datastructures.go"
)

// MergeSort sorts the given []Comparable by merge sort algorithm.
// Time complexity O(nlogn)
// Space complexity O(1)
func MergeSort(c ds.Comparables) ds.Comparables {
	if len(c) <= 1 {
		return c
	}

	mid := int(math.Floor(float64(len(c)) / 2))

	return sortAndMerge(MergeSort(c[:mid]), MergeSort(c[mid:]))
}

func sortAndMerge(a ds.Comparables, b ds.Comparables) ds.Comparables {
	c := make(ds.Comparables, 0)
	an, bn := a.Shift(), b.Shift()

	for an != nil || bn != nil {
		if an == nil {
			c = append(c, bn)
			bn = b.Shift()
		} else if bn == nil {
			c = append(c, an)
			an = a.Shift()
		} else if an.Compare(bn) < 0 {
			c = append(c, an)
			an = a.Shift()
		} else {
			c = append(c, bn)
			bn = b.Shift()
		}
	}

	return c
}

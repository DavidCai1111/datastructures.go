package sort

import (
	"container/heap"

	ds "github.com/DavidCai1993/datastructures.go"
	h "github.com/DavidCai1993/datastructures.go/min-heap"
)

// HeapSort sorts the given []Comparable by heap sort algorithm.
// Time complexity O(nlogn)
// Space complexity O(1)
func HeapSort(c ds.Comparables) ds.Comparables {
	mh := h.MinHeap(c)

	heap.Init(&mh)

	sc := make(ds.Comparables, 0)

	for mh.Len() != 0 {
		sc = append(sc, heap.Pop(&mh).(ds.Comparable))
	}

	return sc
}

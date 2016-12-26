package minheap

import ds "github.com/DavidCai1993/datastructures.go"

// MinHeap represents a binary min-heap which implements
// heap.Interface .
type MinHeap []ds.Comparable

// Len is to implements heap.Interface .
func (bh *MinHeap) Len() int {
	return len(*bh)
}

// Less is to implements heap.Interface .
func (bh *MinHeap) Less(i, j int) bool {
	return (*bh)[i].Compare((*bh)[j]) < 0
}

// Swap is to implements heap.Interface .
func (bh *MinHeap) Swap(i, j int) {
	(*bh)[i], (*bh)[j] = (*bh)[j], (*bh)[i]
}

// Push is to implements heap.Interface .
func (bh *MinHeap) Push(val interface{}) {
	c, ok := val.(ds.Comparable)

	if !ok {
		panic("value's type should be ds.Comparable")
	}

	*bh = append(*bh, c)
}

// Pop is to implements heap.Interface .
func (bh *MinHeap) Pop() (v interface{}) {
	l := len(*bh)
	v = (*bh)[l-1]
	*bh = (*bh)[:l-1]
	return
}

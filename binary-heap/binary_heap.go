package binaryheap

import (
	ds "github.com/DavidCai1993/datastructures.go"
)

// 左儿子 2i
// 右儿子在左儿子后的 (2i + 1) 中
// 父亲 [i / 2]
// 数组第一个元素为空，从小到大排列

// BinaryHeap represents a binary min-heap.
type BinaryHeap []ds.Comparable

// Len is to implements heap.Interface .
func (bh BinaryHeap) Len() int {
	return len(bh)
}

// Less is to implements heap.Interface .
func (bh BinaryHeap) Less(i, j int) bool {
	return bh[i].Compare(bh[j]) < 0
}

// Swap is to implements heap.Interface .
func (bh BinaryHeap) Swap(i, j int) {
	bh[i], bh[j] = bh[j], bh[i]
}

// Push is to implements heap.Interface .
func (bh BinaryHeap) Push(val interface{}) {
	c, ok := val.(ds.Comparable)

	if !ok {
		panic("value's type should be ds.Comparable")
	}

	bh = append(bh, c)
}

// Pop is to implements heap.Interface .
func (bh BinaryHeap) Pop() (v interface{}) {
	l := len(bh)
	v = bh[l-1]
	bh = bh[:l-1]
	return
}

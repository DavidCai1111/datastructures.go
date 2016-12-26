package minheap

import (
	"container/heap"
	"testing"

	ds "github.com/DavidCai1993/datastructures.go"
	"github.com/stretchr/testify/suite"
)

type MinHeapSuite struct {
	suite.Suite

	h *MinHeap
}

func (s *MinHeapSuite) SetupTest() {
	s.h = &MinHeap{
		ds.IntComparable(2),
		ds.IntComparable(1),
		ds.IntComparable(5),
		ds.IntComparable(100),
		ds.IntComparable(3),
		ds.IntComparable(6),
		ds.IntComparable(4),
		ds.IntComparable(5),
	}

	heap.Init(s.h)
}

func (s MinHeapSuite) TestInit() {
	s.Equal(ds.IntComparable(1), (*s.h)[0])
	s.Equal(ds.IntComparable(2), (*s.h)[1])
	s.Equal(ds.IntComparable(3), (*s.h)[4])
}

func (s MinHeapSuite) TestMin() {
	s.Equal(ds.IntComparable(1), heap.Pop(s.h))
	s.Equal(ds.IntComparable(2), heap.Pop(s.h))
	s.Equal(ds.IntComparable(3), heap.Pop(s.h))
	s.Equal(ds.IntComparable(4), heap.Pop(s.h))
	s.Equal(ds.IntComparable(5), heap.Pop(s.h))
	s.Equal(ds.IntComparable(5), heap.Pop(s.h))
	s.Equal(ds.IntComparable(6), heap.Pop(s.h))
	s.Equal(ds.IntComparable(100), heap.Pop(s.h))
}

func TestBinaryHeap(t *testing.T) {
	suite.Run(t, new(MinHeapSuite))
}

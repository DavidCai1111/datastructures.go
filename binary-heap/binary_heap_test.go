package binaryheap

import (
	"container/heap"
	"testing"

	ds "github.com/DavidCai1993/datastructures.go"
	"github.com/stretchr/testify/suite"
)

type BinaryHeapSuite struct {
	suite.Suite

	h BinaryHeap
}

func (s *BinaryHeapSuite) SetupTest() {
	s.h = BinaryHeap{
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

func (s BinaryHeapSuite) TestInit() {
	s.Equal(ds.IntComparable(1), s.h[0])
	s.Equal(ds.IntComparable(2), s.h[1])
	s.Equal(ds.IntComparable(3), s.h[4])
}

func TestBinaryHeap(t *testing.T) {
	suite.Run(t, new(BinaryHeapSuite))
}

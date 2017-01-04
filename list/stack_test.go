package list

import (
	"testing"

	ds "github.com/DavidCai1993/datastructures.go"
	"github.com/stretchr/testify/suite"
)

type StackSuite struct {
	suite.Suite

	s *Stack
}

func (s *StackSuite) SetupTest() {
	s.s = NewStack()
}

func (s StackSuite) TestPush() {
	s.s.Push(ds.IntComparable(10))

	s.Equal(1, s.s.l.Len())
}

func (s StackSuite) TestPop() {
	s.s.Push(ds.IntComparable(1))
	s.s.Push(ds.IntComparable(2))

	s.Equal(ds.IntComparable(2), s.s.Pop())
	s.Equal(ds.IntComparable(1), s.s.Pop())
}

func (s StackSuite) TestTop() {
	s.s.Push(ds.IntComparable(1))
	s.s.Push(ds.IntComparable(2))

	s.Equal(ds.IntComparable(2), s.s.Top())
	s.Equal(ds.IntComparable(2), s.s.Pop())
}

func (s StackSuite) TestEmpty() {
	s.True(s.s.Empty())

	s.s.Push(ds.IntComparable(1))
	s.False(s.s.Empty())
}

func TestStack(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

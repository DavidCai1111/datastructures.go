package datastructures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IntComparableSuite struct {
	suite.Suite
	One   IntComparable
	Two   IntComparable
	Three IntComparable
}

func (s *IntComparableSuite) SetupTest() {
	s.One = IntComparable(1)
	s.Two = IntComparable(2)
	s.Three = IntComparable(3)
}

func (s IntComparableSuite) TestLess() {
	s.True(s.One.Compare(s.Two) < 0)
}

func (s IntComparableSuite) TestMore() {
	s.True(s.Three.Compare(s.Two) > 0)
}

func (s IntComparableSuite) TestEqual() {
	s.Zero((s.One.Compare(s.One)))
}

func (s IntComparableSuite) TestShift() {
	var comparables = Comparables{
		IntComparable(1),
		IntComparable(2),
		IntComparable(3),
	}

	s.Equal(IntComparable(1), comparables.Shift())
	s.Equal(2, len(comparables))
}

func TestIntComparable(t *testing.T) {
	suite.Run(t, new(IntComparableSuite))
}

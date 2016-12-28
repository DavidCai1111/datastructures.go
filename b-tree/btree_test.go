package btree

import (
	"testing"

	ds "github.com/DavidCai1993/datastructures.go"
	"github.com/stretchr/testify/suite"
)

type BTreeSuite struct {
	suite.Suite

	tree *Tree
}

func (s *BTreeSuite) SetupTest() {
	t, err := New(ds.IntComparable(1),
		ds.IntComparable(2),
		ds.IntComparable(3),
		ds.IntComparable(4))

	s.Nil(err)

	s.tree = t
}

func (s BTreeSuite) TestNew() {
	t, err := New(ds.IntComparable(1),
		ds.IntComparable(2),
		ds.IntComparable(3),
		ds.IntComparable(4))

	s.Nil(err)
	s.Equal(ds.IntComparable(3), t.root.lMax)
	s.Equal(nil, t.root.rMax)
	s.Equal(ds.IntComparable(1), t.root.n[0].c[0])
	s.Equal(ds.IntComparable(2), t.root.n[0].c[1])
	s.Equal(ds.IntComparable(3), t.root.n[1].c[0])
	s.Equal(ds.IntComparable(4), t.root.n[1].c[1])
}

func (s BTreeSuite) TestNewLackVals() {
	_, err := New(ds.IntComparable(1))

	s.Equal(ErrLacksInitialVals, err)
}

func TestBTree(t *testing.T) {
	suite.Run(t, new(BTreeSuite))
}

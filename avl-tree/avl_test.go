package avl

import (
	"testing"

	ds "github.com/DavidCai1993/datastructures.go"
	"github.com/stretchr/testify/suite"
)

type AVLSuite struct {
	suite.Suite

	tree *Tree
}

func (s *AVLSuite) SetupTest() {
	tree := New()
	tree.Insert(ds.IntComparable(100))
	tree.Insert(ds.IntComparable(50))
	tree.Insert(ds.IntComparable(150))
	tree.Insert(ds.IntComparable(20))
	tree.Insert(ds.IntComparable(70))
	tree.Insert(ds.IntComparable(200))
	tree.Insert(ds.IntComparable(120))

	s.tree = tree
}

func (s AVLSuite) TestNew() {
	tree := New()

	s.Nil(tree.root.val)
	s.Nil(tree.root.left)
	s.Nil(tree.root.right)
}

func (s AVLSuite) TestFindMin() {
	s.Equal(ds.IntComparable(20), s.tree.FindMin())
}

func (s AVLSuite) TestFindMax() {
	s.Equal(ds.IntComparable(200), s.tree.FindMax())
}

func (s AVLSuite) TestFind() {
	s.True(s.tree.Find(ds.IntComparable(70)))
	s.False(s.tree.Find(ds.IntComparable(-1)))
}

func (s AVLSuite) TestLeftRoation() {
	tree := New()
	tree.Insert(ds.IntComparable(1))
	tree.Insert(ds.IntComparable(2))
	tree.Insert(ds.IntComparable(3))

	s.Equal(ds.IntComparable(2), tree.root.val)
	s.Equal(ds.IntComparable(1), tree.root.left.val)
	s.Equal(ds.IntComparable(3), tree.root.right.val)

	s.Nil(tree.root.left.left)
	s.Nil(tree.root.left.right)
	s.Nil(tree.root.right.left)
	s.Nil(tree.root.right.right)
}

func (s AVLSuite) TestRightRoation() {
	tree := New()
	tree.Insert(ds.IntComparable(3))
	tree.Insert(ds.IntComparable(2))
	tree.Insert(ds.IntComparable(1))

	s.Equal(ds.IntComparable(2), tree.root.val)
	s.Equal(ds.IntComparable(1), tree.root.left.val)
	s.Equal(ds.IntComparable(3), tree.root.right.val)

	s.Nil(tree.root.left.left)
	s.Nil(tree.root.left.right)
	s.Nil(tree.root.right.left)
	s.Nil(tree.root.right.right)
}

func (s AVLSuite) TestLeftRightRoation() {
	tree := New()
	tree.Insert(ds.IntComparable(3))
	tree.Insert(ds.IntComparable(1))
	tree.Insert(ds.IntComparable(2))

	s.Equal(ds.IntComparable(2), tree.root.val)
	s.Equal(ds.IntComparable(1), tree.root.left.val)
	s.Equal(ds.IntComparable(3), tree.root.right.val)

	s.Nil(tree.root.left.left)
	s.Nil(tree.root.left.right)
	s.Nil(tree.root.right.left)
	s.Nil(tree.root.right.right)
}

func (s AVLSuite) TestRightLeftRoation() {
	tree := New()
	tree.Insert(ds.IntComparable(2))
	tree.Insert(ds.IntComparable(3))
	tree.Insert(ds.IntComparable(1))

	s.Equal(ds.IntComparable(2), tree.root.val)
	s.Equal(ds.IntComparable(1), tree.root.left.val)
	s.Equal(ds.IntComparable(3), tree.root.right.val)

	s.Nil(tree.root.left.left)
	s.Nil(tree.root.left.right)
	s.Nil(tree.root.right.left)
	s.Nil(tree.root.right.right)
}

func (s AVLSuite) TestGetHeight() {
	tree := New()
	tree.root.left = &node{val: ds.IntComparable(1)}
	tree.root.left.left = &node{val: ds.IntComparable(1)}
	tree.root.left.right = &node{val: ds.IntComparable(1)}
	tree.root.left.left.left = &node{val: ds.IntComparable(1)}

	s.Equal(3, tree.root.getHeight(0))
}

func (s AVLSuite) TestSelfBlance() {
	s.tree.Insert(ds.IntComparable(10))
	s.tree.Insert(ds.IntComparable(15))

	s.Equal(3, s.tree.root.left.getHeight(1))
	s.Equal(2, s.tree.root.right.getHeight(1))
}

func TestAVL(t *testing.T) {
	suite.Run(t, new(AVLSuite))
}

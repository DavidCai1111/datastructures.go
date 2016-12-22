package bst

import (
	"testing"

	ds "github.com/DavidCai1993/datastructures.go"
	"github.com/stretchr/testify/suite"
)

type BSTSuite struct {
	suite.Suite

	tree *Tree
}

func (s *BSTSuite) SetupTest() {
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

func (s BSTSuite) TestNew() {
	tree := New()

	s.Nil(tree.root.val)
	s.Nil(tree.root.left)
	s.Nil(tree.root.right)
}

func (s BSTSuite) TestInsert() {
	s.Equal(ds.IntComparable(100), s.tree.root.val)
	s.Equal(ds.IntComparable(50), s.tree.root.left.val)
	s.Equal(ds.IntComparable(150), s.tree.root.right.val)
	s.Equal(ds.IntComparable(20), s.tree.root.left.left.val)
	s.Equal(ds.IntComparable(70), s.tree.root.left.right.val)
	s.Equal(ds.IntComparable(200), s.tree.root.right.right.val)
	s.Equal(ds.IntComparable(120), s.tree.root.right.left.val)
}

func (s BSTSuite) TestFindMin() {
	s.Equal(ds.IntComparable(20), s.tree.FindMin())
}

func (s BSTSuite) TestFindMax() {
	s.Equal(ds.IntComparable(200), s.tree.FindMax())
}

func (s BSTSuite) TestFind() {
	s.True(s.tree.Find(ds.IntComparable(70)))
	s.False(s.tree.Find(ds.IntComparable(-1)))
}

func (s BSTSuite) TestDeleteMinNode() {
	s.NotNil(s.tree.root.left.left)

	s.True(s.tree.Find(ds.IntComparable(20)))
	s.tree.Delete(ds.IntComparable(20))
	s.False(s.tree.Find(ds.IntComparable(20)))

	s.Nil(s.tree.root.left.left)
}

func (s BSTSuite) TestDeleteOneChildNode() {
	s.tree.Insert(ds.IntComparable(10))
	s.tree.Delete(ds.IntComparable(20))

	s.Equal(ds.IntComparable(10), s.tree.root.left.left.val)
}

func (s BSTSuite) TestDeleteTwoChildrenNodeWithRighLeftIsNil() {
	s.tree.Delete(ds.IntComparable(150))

	s.Equal(ds.IntComparable(200), s.tree.root.right.val)
	s.Nil(s.tree.root.right.right)
	s.Equal(ds.IntComparable(120), s.tree.root.right.left.val)
}

func (s BSTSuite) TestDeleteTwoChildrenNodeWithRighLeftIsNotNil() {
	s.tree.Insert(ds.IntComparable(170))
	s.tree.Insert(ds.IntComparable(250))

	s.tree.Delete(ds.IntComparable(150))

	s.Equal(ds.IntComparable(170), s.tree.root.right.val)
	s.Equal(ds.IntComparable(120), s.tree.root.right.left.val)
	s.Equal(ds.IntComparable(200), s.tree.root.right.right.val)
	s.Equal(ds.IntComparable(250), s.tree.root.right.right.right.val)
}

func TestBST(t *testing.T) {
	suite.Run(t, new(BSTSuite))
}

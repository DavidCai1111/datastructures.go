package pat

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PATSuite struct {
	suite.Suite

	tree *Tree
}

func (s PATSuite) TestNew() {
	tree := New("XMADAMYX")

	s.Empty(tree.root.val)

	s.Equal(6, len(tree.root.children))
	s.Equal("X", tree.root.children[0].val)
	s.Equal(2, len(tree.root.children[0].children))
	s.Equal("MADAMYX$", tree.root.children[0].children[0].val)
	s.Equal("$", tree.root.children[0].children[1].val)
	s.Equal("M", tree.root.children[1].val)
	s.Equal(2, len(tree.root.children[1].children))
	s.Equal("ADAMYX$", tree.root.children[1].children[0].val)
	s.Equal("YX$", tree.root.children[1].children[1].val)
	s.Equal("A", tree.root.children[2].val)
	s.Equal(2, len(tree.root.children[2].children))
	s.Equal("DAMYX$", tree.root.children[2].children[0].val)
	s.Equal("MYX$", tree.root.children[2].children[1].val)
	s.Equal("DAMYX$", tree.root.children[3].val)
	s.Equal("YX$", tree.root.children[4].val)
	s.Equal("$", tree.root.children[5].val)
}

func TestPAT(t *testing.T) {
	suite.Run(t, new(PATSuite))
}

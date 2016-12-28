package pat

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PATSuite struct {
	suite.Suite

	tree *Tree
}

func (s *PATSuite) SetupTest() {
	s.tree = New("XMADAMYX")
}

func (s PATSuite) TestNew() {
	s.Empty(s.tree.root.val)

	s.Equal(6, len(s.tree.root.children))
	s.Equal("X", s.tree.root.children[0].val)
	s.Equal(2, len(s.tree.root.children[0].children))
	s.Equal("MADAMYX$", s.tree.root.children[0].children[0].val)
	s.Equal("$", s.tree.root.children[0].children[1].val)
	s.Equal("M", s.tree.root.children[1].val)
	s.Equal(2, len(s.tree.root.children[1].children))
	s.Equal("ADAMYX$", s.tree.root.children[1].children[0].val)
	s.Equal("YX$", s.tree.root.children[1].children[1].val)
	s.Equal("A", s.tree.root.children[2].val)
	s.Equal(2, len(s.tree.root.children[2].children))
	s.Equal("DAMYX$", s.tree.root.children[2].children[0].val)
	s.Equal("MYX$", s.tree.root.children[2].children[1].val)
	s.Equal("DAMYX$", s.tree.root.children[3].val)
	s.Equal("YX$", s.tree.root.children[4].val)
	s.Equal("$", s.tree.root.children[5].val)
}

func (s PATSuite) TestNewFailed() {
	s.Panics(func() { New("") })
	s.Panics(func() { New("ABC$") })
}

func (s PATSuite) TestHasSuffix() {
	s.True(s.tree.HasSuffix("DAMYX"))
	s.False(s.tree.HasSuffix("DDAMYX"))
	s.True(s.tree.HasSuffix("XMADAMYX"))
}

func (s PATSuite) TestHasSuffixFailed() {
	s.Panics(func() { s.tree.HasSuffix("") })
	s.Panics(func() { s.tree.HasSuffix("X$") })
}

func TestPAT(t *testing.T) {
	suite.Run(t, new(PATSuite))
}

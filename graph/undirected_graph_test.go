package graph

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UndirectedGraphSuite struct {
	suite.Suite

	g *UndirectedGraph
}

func (s *UndirectedGraphSuite) SetupTest() {
	s.g = New(10)
}

func (s UndirectedGraphSuite) TestNew() {
	s.Equal(10, len(s.g.bags))
}

func (s UndirectedGraphSuite) TestVertexs() {
	s.Equal(10, s.g.Vertexs())
}

func (s UndirectedGraphSuite) TestEdgesFailed() {
	s.Panics(func() { s.g.addEdge(11, 1) })
	s.Panics(func() { s.g.addEdge(1, 11) })
}

func (s UndirectedGraphSuite) TestGetAdjacentVertexs() {
	s.g.addEdge(0, 1)
	s.g.addEdge(0, 2)

	s.Equal([]int{1, 2}, s.g.GetAdjacentVertexs(0))
}

func (s UndirectedGraphSuite) TestDegree() {
	s.g.addEdge(0, 1)
	s.g.addEdge(0, 2)

	s.Equal(2, s.g.Degree(0))
}

func (s UndirectedGraphSuite) TestMaxDegree() {
	s.g.addEdge(0, 1)
	s.g.addEdge(0, 2)
	s.g.addEdge(1, 2)

	s.Equal(2, s.g.MaxDegree())
}

func (s UndirectedGraphSuite) TestPathTo() {
	s.g.addEdge(0, 1)
	s.g.addEdge(1, 2)
	s.g.addEdge(1, 3)
	s.g.addEdge(2, 3)
	s.g.addEdge(4, 5)

	p, ok := s.g.PathTo(0, 2)

	s.Equal([]int{0, 1, 2}, p)
	s.True(ok)

	_, ok = s.g.PathTo(0, 5)

	s.False(ok)
}

func (s UndirectedGraphSuite) TestHasPath() {
	s.g.addEdge(0, 1)
	s.g.addEdge(0, 2)
	s.g.addEdge(1, 2)
	s.g.addEdge(1, 3)
	s.g.addEdge(2, 3)

	s.True(s.g.HasPathTo(0, 2))
}

func TestUndirectedGraph(t *testing.T) {
	suite.Run(t, new(UndirectedGraphSuite))
}

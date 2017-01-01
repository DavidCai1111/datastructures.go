package graph

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GraphSuite struct {
	suite.Suite

	g *Graph
}

func (s *GraphSuite) SetupTest() {
	s.g = New(10)
}

func (s GraphSuite) TestNew() {
	s.Equal(10, len(s.g.bags))
}

func (s GraphSuite) TestVertexs() {
	s.Equal(10, s.g.Vertexs())
}

func (s GraphSuite) TestEdges() {
	s.Equal(0, s.g.Edges())
	s.g.addEdge(0, 1)
	s.Equal(1, s.g.Edges())
}

func (s GraphSuite) TestEdgesFailed() {
	s.Panics(func() { s.g.addEdge(11, 1) })
	s.Panics(func() { s.g.addEdge(1, 11) })
}

func (s GraphSuite) TestGetAdjacentVertexs() {
	s.g.addEdge(0, 1)
	s.g.addEdge(0, 2)

	s.Equal([]int{1, 2}, s.g.GetAdjacentVertexs(0))
}

func TestGraph(t *testing.T) {
	suite.Run(t, new(GraphSuite))
}

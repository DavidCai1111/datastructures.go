package graph

import (
	"fmt"
)

const errInvaliVertexNumber = "graph: invalid vertex number %v"

// UndirectedGraph represents an undirected graph.
type UndirectedGraph struct {
	bags  [][]*vertex
	edges int
}

type vertex struct {
	i         int
	isVisited bool
}

func (v vertex) findPathTo(s []int, w int) ([]int, bool) {
	return nil, false
}

// New returns a new undirected graph instance with v vertexs.
func New(v int) *UndirectedGraph {
	return &UndirectedGraph{bags: make([][]*vertex, v)}
}

// Vertexs returns the number of vertexs of the graph.
func (g UndirectedGraph) Vertexs() int {
	return len(g.bags)
}

// Edges returns the number of edges of the graph.
func (g UndirectedGraph) Edges() int {
	return g.edges
}

// addEdge adds an edge to the graph.
func (g *UndirectedGraph) addEdge(v, w int) {
	if v < 0 || v >= len(g.bags) {
		panic(fmt.Errorf(errInvaliVertexNumber, v))
	}

	if w < 0 || w >= len(g.bags) {
		panic(fmt.Errorf(errInvaliVertexNumber, w))
	}

	for _, b := range g.bags[v] {
		if b.i == w {
			panic(fmt.Errorf(errInvaliVertexNumber, w))
		}
	}

	g.bags[v] = append(g.bags[v], &vertex{i: w})

	for _, b := range g.bags[w] {
		if b.i == v {
			panic(fmt.Errorf(errInvaliVertexNumber, v))
		}
	}

	g.bags[w] = append(g.bags[w], &vertex{i: v})

	g.edges++
}

// GetAdjacentVertexs returns all adjacent vertexs of v in the graph.
func (g *UndirectedGraph) GetAdjacentVertexs(v int) []int {
	if v < 0 || v >= len(g.bags) {
		panic(fmt.Errorf(errInvaliVertexNumber, v))
	}

	vs := make([]int, len(g.bags[v]))

	for i, ver := range g.bags[v] {
		vs[i] = ver.i
	}

	return vs
}

// Degree returns the degree of the vertexs v.
func (g *UndirectedGraph) Degree(v int) int {
	return len(g.GetAdjacentVertexs(v))
}

// MaxDegree returns the max degree of the graph.
func (g *UndirectedGraph) MaxDegree() int {
	max := 0

	for _, adj := range g.bags {
		if len(adj) > max {
			max = len(adj)
		}
	}

	return max
}

// PathTo returns path from vertex v to w using breadth first traversal.
func (g *UndirectedGraph) PathTo(v, w int) ([]int, bool) {
	if v < 0 || v >= len(g.bags) {
		panic(fmt.Errorf(errInvaliVertexNumber, v))
	}

	if w < 0 || w >= len(g.bags) {
		panic(fmt.Errorf(errInvaliVertexNumber, w))
	}

	if s, ok := g.pathTo(v, w, []int{v}); ok {
		return s, true
	}

	return nil, false
}

func (g *UndirectedGraph) pathTo(v, w int, s []int) ([]int, bool) {
	for _, n := range g.bags[v] {
		if n.isVisited {
			break
		}

		if n.i == w {
			return append(s, w), true
		}

		n.isVisited = true

		if p, ok := g.pathTo(n.i, w, append(s, n.i)); ok {
			return p, true
		}
	}

	return nil, false
}

// HasPathTo checks whether the graph has a path from vertex v to w.
func (g *UndirectedGraph) HasPathTo(v, w int) bool {
	if _, ok := g.PathTo(v, w); ok {
		return true
	}

	return false
}

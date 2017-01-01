package graph

import (
	"fmt"
)

const errInvaliVertexNumber = "graph: invalid vertex number %v"

// Graph represents an undirected graph.
type Graph struct {
	bags  [][]int
	edges int
}

// New returns a new undirected graph instance with v vertexs.
func New(v int) *Graph {
	return &Graph{bags: make([][]int, v)}
}

// Vertexs returns the number of vertexs of the graph.
func (g Graph) Vertexs() int {
	return len(g.bags)
}

// Edges returns the number of edges of the graph.
func (g Graph) Edges() int {
	return g.edges
}

// addEdge adds an edge to the graph.
func (g *Graph) addEdge(v, w int) {
	if v < 0 || v >= len(g.bags) {
		panic(fmt.Errorf(errInvaliVertexNumber, v))
	}

	if w < 0 || w >= len(g.bags) {
		panic(fmt.Errorf(errInvaliVertexNumber, w))
	}

	for _, b := range g.bags[v] {
		if b == w {
			panic(fmt.Errorf(errInvaliVertexNumber, w))
		}
	}

	g.bags[v] = append(g.bags[v], w)

	for _, b := range g.bags[w] {
		if b == v {
			panic(fmt.Errorf(errInvaliVertexNumber, v))
		}
	}

	g.bags[w] = append(g.bags[w], v)

	g.edges++
}

// GetAdjacentVertexs returns all adjacent vertexs of v in the graph.
func (g *Graph) GetAdjacentVertexs(v int) []int {
	if v < 0 || v >= len(g.bags) {
		panic(fmt.Errorf(errInvaliVertexNumber, v))
	}

	return g.bags[v]
}

package btree

import (
	"errors"
	"sort"

	ds "github.com/DavidCai1993/datastructures.go"
)

// Errors used by this package
var (
	ErrLacksInitialVals = errors.New("bst: number of initial values is less than 2")
)

// Tree represents a 2-3 tree.
type Tree struct {
	root *node
}

type (
	children [3]ds.Comparable
	nodes    [3]*node
)

type node struct {
	lMax ds.Comparable
	rMax ds.Comparable
	n    nodes
	c    *children
}

// New returns a new instance 2-3 tree. Because the tree needs
// each node at least has 2 children, so we need at least 4 initial values.
func New(vals ds.Comparables) (*Tree, error) {
	if len(vals) < 4 {
		return nil, ErrLacksInitialVals
	}

	n := &node{}
	t := &Tree{root: n}
	sort.Sort(vals)

	n.lMax = vals[2]

	n.n[0].c[0] = vals[0]
	n.n[0].c[1] = vals[1]
	n.n[1].c[0] = vals[2]
	n.n[1].c[1] = vals[3]

	for _, val := range vals[3:] {
		t.Insert(val)
	}

	return t, nil
}

// Insert inserts a value to the tree.
func (t *Tree) Insert(val ds.Comparable) {

}

// Find checks whether the given value is this tree.
func (t *Tree) Find(val ds.Comparable) bool {
	return t.root.find(val)
}

func (n *node) find(val ds.Comparable) bool {
	if n.c != nil {
		for _, c := range n.c {
			if c.Compare(val) == 0 {
				return true
			}
		}

		return false
	}

	if n.lMax == nil || val.Compare(n.lMax) < 0 {
		return n.n[0].find(val)
	}

	if n.rMax != nil {
		if val.Compare(n.lMax) >= 0 || val.Compare(n.rMax) < 0 {
			return n.n[1].find(val)
		}

		return n.n[2].find(val)
	}

	return false
}

func (ns nodes) isFull() bool {
	return ns[2] != nil
}

func (c children) isFull() bool {
	return c[2] != nil
}

package btree

import (
	"errors"
	"sort"

	ds "github.com/DavidCai1993/datastructures.go"
)

// Errors used by this package
var (
	ErrLacksInitialVals = errors.New("bst: number of initial values is less than 4")
)

type children []ds.Comparable

func (c children) has(val ds.Comparable) bool {
	for _, v := range c {
		if v.Compare(val) == 0 {
			return true
		}
	}

	return false
}

func (c children) insertInOrder(val ds.Comparable) {
	c = append(c, val)
	sort.Sort(ds.Comparables(c))
}

// Tree represents a 2-3 tree.
type Tree struct {
	root *node
}

type node struct {
	lMax   ds.Comparable
	rMax   ds.Comparable
	parent *node
	n      []*node
	c      children
}

// New returns a new instance 2-3 tree. Because the tree needs
// each node at least has 2 children, so we need at least 4 initial values.
func New(vals ...ds.Comparable) (*Tree, error) {
	if len(vals) < 4 {
		return nil, ErrLacksInitialVals
	}

	sort.Sort(ds.Comparables(vals))

	t := &Tree{
		root: &node{lMax: vals[2], n: []*node{
			&node{c: children{vals[0], vals[1]}},
			&node{c: children{vals[2], vals[3]}},
		}},
	}

	for _, val := range vals[4:] {
		t.Insert(val)
	}

	return t, nil
}

// Insert inserts a value to the tree.
func (t *Tree) Insert(val ds.Comparable) {
	t.root.insert(val)
}

func (n *node) insert(val ds.Comparable) {
	f := n.findNode(val)

	if f.c.has(val) {
		return
	}

	if len(f.c) < 3 {
		f.c.insertInOrder(val)
		return
	}

	if len(f.parent.n) < 3 {
		f.parent.reshuffleWithVal(val)
		return
	}

	if f.parent.childrenCount() < 9 &&
		val.Compare(f.parent.getMin()) > 0 &&
		val.Compare(f.parent.getMax()) < 0 {
		f.parent.reshuffleWithVal(val)
		return
	}

	if len(f.parent.parent.n) < 3 {
		f.parent.parent.addNodeWithVal(val)
		return
	}

	f.parent.parent.addHeightWithVal(val)
}

func (n *node) getMin() ds.Comparable {
	return n.n[0].c[0]
}

func (n *node) getMax() ds.Comparable {
	if len(n.n) == 2 {
		return n.n[1].c[len(n.n[1].c)-1]
	}

	return n.n[2].c[len(n.n[2].c)-1]
}

func (n *node) childrenCount() int {
	count := 0

	for _, nod := range n.n {
		count += len(nod.c)
	}

	return count
}

func (n *node) getVals() []ds.Comparable {
	vals := []ds.Comparable{}

	for _, nod := range n.n {
		for _, c := range nod.c {
			vals = append(vals, c)
		}
	}

	return vals
}

func (n *node) reshuffleWithVal(val ds.Comparable) {
	vals := append(n.getVals(), val)

	sort.Sort(ds.Comparables(vals))

	n.lMax = vals[2]
	n.rMax = vals[4]

	n.n = []*node{
		&node{c: children{vals[0], vals[1]}},
		&node{c: children{vals[2], vals[3]}},
		&node{c: children{vals[4], vals[5], vals[6]}},
	}
}

func (n *node) addNodeWithVal(val ds.Comparable) {
	vals := []ds.Comparable{}

	for _, nod := range n.n {
		vals = append(vals, nod.getVals()...)
	}

	sort.Sort(ds.Comparables(vals))
}

func (n *node) addHeightWithVal(val ds.Comparable) {
}

// Find checks whether the given value is this tree.
func (t *Tree) Find(val ds.Comparable) bool {
	return t.root.find(val)
}

func (n *node) findNode(val ds.Comparable) *node {
	if n.c != nil && len(n.c) > 0 {
		return n
	}

	if val.Compare(n.lMax) < 0 {
		return n.n[0].findNode(val)
	}

	if n.rMax == nil || val.Compare(n.rMax) < 0 {
		return n.n[1].findNode(val)
	}

	return n.n[2].findNode(val)
}

func (n node) find(val ds.Comparable) bool {
	for _, c := range n.findNode(val).c {
		if c.Compare(val) == 0 {
			return true
		}
	}

	return false
}

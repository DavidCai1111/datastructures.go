package avl

import (
	"errors"
	"fmt"
	"math"

	ds "github.com/DavidCai1993/datastructures.go"
)

// Errors used by this package
var (
	ErrEmptyTree = errors.New("avl: empty tree")
)

// Tree represents an AVL (Adelson, Velski & Landis) tree.
type Tree struct {
	root *node
}

type node struct {
	val    ds.Comparable
	parent *node
	left   *node
	right  *node
}

// New returns an new AVL tree instance.
func New() *Tree {
	return &Tree{root: &node{}}
}

// Insert inserts a value to the tree.
// Time Complexity: O(logN)
func (t *Tree) Insert(val ds.Comparable) {
	n := t.root.insert(val, nil)

	// an unbalance tree should at least have a height of 2 .
	if n.parent == nil || n.parent.parent == nil {
		return
	}

	n.parent.parent.ensureBalance()
}

func (n *node) insert(val ds.Comparable, parent *node) *node {
	if n.val == nil {
		n.val = val
		n.parent = parent
		return n
	}

	c := val.Compare(n.val)

	if c == 0 {
		return n
	} else if c < 0 {
		n.left = ensureNode(n.left)
		return n.left.insert(val, n)
	} else {
		n.right = ensureNode(n.right)
		return n.right.insert(val, n)
	}
}

func (n *node) ensureBalance() {
	lh, rh := 0, 0

	if !n.left.isNil() {
		lh = n.left.getHeight(1)
	}

	if !n.right.isNil() {
		rh = n.right.getHeight(1)
	}

	diff := int(math.Abs(float64(lh - rh)))

	if diff == 1 || diff == 0 {
		return
	}

	if diff != 2 {
		panic(fmt.Errorf("invalid diff: %v", diff))
	}

	// need left rotation
	if n.left.isNil() && !n.right.isNil() && !n.right.right.isNil() {
		n.right.left = &node{val: n.val}
		*n = *n.right
		return
	}

	// need right rotation
	if n.right.isNil() && !n.left.isNil() && !n.left.left.isNil() {
		n.left.right = &node{val: n.val}
		*n = *n.left
		return
	}

	// need left-right rotation
	if n.right.isNil() && !n.left.isNil() && !n.left.right.isNil() {
		n.left.right.left = &node{val: n.left.val}
		*n.left = *n.left.right
		n.left.right = &node{val: n.val}
		*n = *n.left
		return
	}

	// need right-left rotation
	if n.left.isNil() && !n.right.isNil() && !n.right.left.isNil() {
		n.right.left.right = &node{val: n.right.val}
		*n.right = *n.right.left
		n.right.left = &node{val: n.val}
		*n = *n.right
		return
	}

	panic("invalid unbalanced situation")
}

func (n *node) isNil() bool {
	return n == nil || n.val == nil
}

func (n *node) getHeight(h int) int {
	hl, hr := h, h

	if !n.left.isNil() {
		hl = n.left.getHeight(h + 1)
	}

	if !n.right.isNil() {
		hr = n.right.getHeight(h + 1)
	}

	if hl > hr {
		return hl
	}

	return hr
}

// FindMin returns the value of the minimum node.
// Time Complexity: O(logN)
func (t Tree) FindMin() ds.Comparable {
	return t.root.findMin()
}

func (n node) findMin() ds.Comparable {
	if n.left == nil {
		return n.val
	}

	return n.left.findMin()
}

// FindMax returns the value of the maximum node.
// Time Complexity: O(logN)
func (t Tree) FindMax() ds.Comparable {
	return t.root.findMax()
}

func (n node) findMax() ds.Comparable {
	if n.right == nil {
		return n.val
	}

	return n.right.findMax()
}

// Find checks whether the given value is this tree.
// Time Complexity: O(logN)
func (t Tree) Find(val ds.Comparable) bool {
	n := t.root.find(val)

	if n != nil {
		return true
	}

	return false
}

func (n *node) find(val ds.Comparable) *node {
	c := val.Compare(n.val)

	if c == 0 {
		return n
	}

	if c < 0 {
		if n.left != nil {
			return n.left.find(val)
		}
	} else {
		if n.right != nil {
			return n.right.find(val)
		}
	}

	return nil
}

func ensureNode(n *node) *node {
	if n == nil {
		n = &node{}
	}

	return n
}

package bst

import (
	"errors"

	ds "github.com/DavidCai1993/datastructures.go"
)

// Errors used by this package
var (
	ErrEmptyTree = errors.New("bst: empty tree")
)

type node struct {
	val    ds.Comparable
	parent *node
	left   *node
	right  *node
}

// Tree represents a binary search tree.
type Tree struct {
	root *node
}

// New returns a new binary search tree instance.
func New() *Tree {
	return &Tree{root: &node{}}
}

// Insert inserts a value to the tree.
// Time Complexity: O(logN)
func (t *Tree) Insert(val ds.Comparable) {
	t.root.insert(val, nil)
}

func (n *node) insert(val ds.Comparable, parent *node) {
	if n.val == nil {
		n.val = val
		n.parent = parent
		return
	}

	c := val.Compare(n.val)

	if c == 0 {
		return
	} else if c < 0 {
		n.left = ensureNode(n.left)
		n.left.insert(val, n)
	} else {
		n.right = ensureNode(n.right)
		n.right.insert(val, n)
	}
}

// Delete removes the node of given value in this tree.
// Time Complexity: O(logN)
func (t *Tree) Delete(val ds.Comparable) {
	n := t.root.find(val)

	if n == nil {
		return
	}

	n.delete()
}

func (n *node) delete() {
	if n.left == nil && n.right == nil {
		if n.parent.left == n {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
		return
	}

	if n.left != nil && n.right == nil {
		*n = *n.left
		return
	}

	if n.left == nil && n.right != nil {
		*n = *n.right
		return
	}

	if n.right.left == nil {
		n.val = n.right.val
		n.right.delete()
		return
	}

	n.val = n.right.left.val

	n.right.left.delete()
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

// PreOrderTraversal runs function f of each node's value in the tree
// by pre order traversal.
func (t Tree) PreOrderTraversal(f func(ds.Comparable)) {
	t.root.preOrderTraversal(f)
}

func (n node) preOrderTraversal(f func(ds.Comparable)) {
	if n.val == nil {
		return
	}

	f(n.val)

	if n.left != nil {
		n.left.preOrderTraversal(f)
	}

	if n.right != nil {
		n.right.preOrderTraversal(f)
	}
}

// InOrderTraversal runs function f of each node's value in the tree
// by in order traversal.
func (t Tree) InOrderTraversal(f func(ds.Comparable)) {
	t.root.inOrderTraversal(f)
}

func (n node) inOrderTraversal(f func(ds.Comparable)) {
	if n.val == nil {
		return
	}

	if n.left != nil {
		n.left.inOrderTraversal(f)
	}

	f(n.val)

	if n.right != nil {
		n.right.inOrderTraversal(f)
	}
}

// PostOrderTraversal runs function f of each node's value in the tree
// by post order traversal.
func (t Tree) PostOrderTraversal(f func(ds.Comparable)) {
	t.root.postOrderTraversal(f)
}

func (n node) postOrderTraversal(f func(ds.Comparable)) {
	if n.val == nil {
		return
	}

	if n.left != nil {
		n.left.postOrderTraversal(f)
	}

	if n.right != nil {
		n.right.postOrderTraversal(f)
	}

	f(n.val)
}

func ensureNode(n *node) *node {
	if n == nil {
		n = &node{}
	}

	return n
}

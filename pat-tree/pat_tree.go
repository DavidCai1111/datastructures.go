package pat

import (
	"fmt"
	"strings"
)

type node struct {
	val      string
	children []*node
}

// Tree represents a PAT tree.
type Tree struct {
	root *node
}

// New returns a new PAT tree instance.
func New(s string) *Tree {
	t := &Tree{root: &node{val: ""}}

	t.insert(s)

	return t
}

func (t *Tree) insert(s string) {
	if strings.HasSuffix(s, "$") {
		panic(fmt.Errorf("string: %v should not have the sufix of '$'", s))
	}

	if len(s) == 0 {
		panic(fmt.Errorf("s is a empty string"))
	}

	s += "$"

	for i := 0; i < len(s); i++ {
		t.root.insert(s[i:])
	}

	for _, c := range t.root.children {
		c.compress()
	}
}

func (n *node) insert(s string) {
loop:
	for _, b := range s {
		str := string(b)
		for _, c := range n.children {
			if c.val == str {
				n = c
				continue loop
			}
		}

		c := &node{val: str}
		n.children = append(n.children, c)
		n = c
	}
}

func (n *node) compress() {
	if n.children == nil {
		return
	}

	if len(n.children) == 1 {
		n.val += n.children[0].val
		n.children = n.children[0].children

		n.compress()
		return
	}

	for _, c := range n.children {
		c.compress()
	}
}

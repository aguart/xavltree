package xavltree

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// node object
type node struct {
	left   *node
	right  *node
	key    uint64
	value  interface{}
	height int
}

// Tree object
type Tree struct {
	root  *node
	count int
}

// NewIntKeys create new empty tree with int keys
func NewTree() *Tree {
	return &Tree{}
}

// TestElements create test tree
func (t *Tree) TestElements(count int) {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		t.Add(uint64(r.Intn(9999)), i)
	}
}

// Add new node to the tree
func (t *Tree) Add(key uint64, value interface{}) (err error) {
	t.root = t.root.add(key, value)
	if err != nil {
		return err
	}
	t.count++
	return nil
}

// Remove node from tree by key
func (t *Tree) Remove(key uint64) (error, bool) {
	if t.root == nil {
		return errors.New("Cannot remove from an empty tree"), false
	}
	_, ok := t.root.remove(key)
	if ok {
		t.count--
	}
	return nil, ok
}

// Get element from tree
func (t *Tree) Get(key uint64) (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.root.get(key)
}

// Count return number of elements
func (t *Tree) Count() int {
	return t.count
}

// Min returm min element from the tree
func (t *Tree) Min() (uint64, interface{}, error) {
	if t.root == nil {
		return 0, nil, errors.New("empty tree")
	}
	key, val := t.root.min()
	return key, val, nil
}

// Max returm max element from the tree
func (t *Tree) Max() (uint64, interface{}, error) {
	if t.root == nil {
		return 0, nil, errors.New("empty tree")
	}
	key, val := t.root.max()
	return key, val, nil
}

// PrintTree internal recursive function to print a tree
func PrintTree(n *node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		PrintTree(n.left, level)
		fmt.Printf(format+"%v\n", n.key)
		PrintTree(n.right, level)
	}
}

func Traverse(n *node, f func(*node)) {
	if n == nil {
		return
	}
	Traverse(n.left, f)
	f(n)
	Traverse(n.right, f)
}

// ==== local funcs ============================================================

// add local method
func (n *node) add(key uint64, value interface{}) *node {
	if n == nil {
		return &node{nil, nil, key, value, 1}
	}
	if key > n.key {
		n.right = n.right.add(key, value)
	} else if key < n.key {
		n.left = n.left.add(key, value)
	} else {
		n.key = key
	}
	return n.rebalanceTree()
}

// findMin -
func (n *node) findMin() *node {
	if n.left != nil {
		return n.left.findMin()
	} else {
		return n
	}
}

// remove -
func (n *node) remove(key uint64) (*node, bool) {
	exist := false
	if n == nil {
		return nil, exist
	}

	if key < n.key {
		n.left, exist = n.left.remove(key)
	} else if key > n.key {
		n.right, exist = n.right.remove(key)
	} else {
		exist = true
		if n.left != nil && n.right != nil {
			rightMinNode := n.right.findMin()
			n.key = rightMinNode.key
			n.right, exist = n.right.remove(rightMinNode.key)
		} else if n.left != nil {
			n = n.left
		} else if n.right != nil {
			n = n.right
		} else {
			n = nil
			return n, exist
		}
	}
	return n.rebalanceTree(), exist
}

// get -
func (n *node) get(key uint64) (interface{}, bool) {
	if n == nil {
		return 0, false
	}
	if key > n.key {
		return n.right.get(key)
	} else if key < n.key {
		return n.left.get(key)
	} else {
		return n.value, true
	}
}

// max -
func (n *node) max() (uint64, interface{}) {
	if n.right == nil {
		return n.key, n.value
	}
	return n.right.max()
}

// min -
func (n *node) min() (uint64, interface{}) {
	if n.left == nil {
		return n.key, n.value
	}
	return n.left.min()
}

// === methods for AVL balansing ===============================================

// getHeight -
func (n *node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

// recalculateHeight -
func (n *node) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

// max -
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// rotateLeft rotate nodes left to balance node
func (n *node) rotateLeft() *node {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// rotateRight rotate nodes right to balance node
func (n *node) rotateRight() *node {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// rebalanceTree checks if node is balanced and rebalance
func (n *node) rebalanceTree() *node {
	if n == nil {
		return n
	}
	n.recalculateHeight()
	// check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		// check if child is left-heavy and rotateRight first
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		// check if child is right-heavy and rotateLeft first
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

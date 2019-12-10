package xavltree

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// node object
type node struct {
	left   *node
	right  *node
	key    interface{}
	value  interface{}
	height int
}

// Tree object
type Tree struct {
	root  *node
	cmp   comparator
	count int
}

// NewIntKeys create new empty tree with int keys
func NewIntKeys() *Tree {
	return &Tree{
		cmp: intComparator,
	}
}

// NewInt64Keys create new empty tree with int64 keys
func NewInt64Keys() *Tree {
	return &Tree{
		cmp: int64Comparator,
	}
}

// NewStrKeys create new empty tree with float64 keys
func NewFloat64Keys() *Tree {
	return &Tree{
		cmp: float64Comparator,
	}
}

// NewStrKeys create new empty tree with string keys
func NewStrKeys() *Tree {
	return &Tree{
		cmp: stringComparator,
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// TestElements create test tree
func (t *Tree) TestElements(count int) {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	rawfName := getFunctionName(t.cmp)
	arr := strings.Split(rawfName, ".")
	switch arr[len(arr)-1] {
	case "intComparator":
		for i := 0; i < count; i++ {
			t.Add(r.Intn(9999), i)
		}
	case "int64Comparator":
		for i := 0; i < count; i++ {
			t.Add(r.Int63n(9999999999999999), i)
		}
	case "float64Comparator":
		for i := 0; i < count; i++ {
			t.Add(r.Float64()*9999, i)
		}
	case "stringComparator":
		letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for i := 0; i < count; i++ {
			b := make([]byte, 8)
			for i := range b {
				b[i] = letterBytes[r.Int63()%int64(len(letterBytes))]
			}
			t.Add(string(b), i)
		}
	}

}

// Add new node to the tree
func (t *Tree) Add(key interface{}, value interface{}) (err error) {
	t.root, err = t.root.add(key, value, t.cmp)
	if err != nil {
		return err
	}
	t.count++
	return nil
}

// Remove node from tree by key
func (t *Tree) Remove(key interface{}) (error, bool) {
	if t.root == nil {
		return errors.New("Cannot remove from an empty tree"), false
	}
	_, ok, err := t.root.remove(key, t.cmp)
	if err != nil {
		return err, false
	}
	if ok {
		t.count--
	}
	return nil, ok
}

func (t *Tree) Get(key interface{}) (interface{}, bool, error) {
	if t.root == nil {
		return nil, false, errors.New("Cannot get from an empty tree")
	}
	return t.root.get(key, t.cmp)
}

// Count return number of elements
func (t *Tree) Count() int {
	return t.count
}

// Min returm min element from the tree
func (t *Tree) Min() (interface{}, interface{}, error) {
	if t.root == nil {
		return 0, nil, errors.New("empty tree")
	}
	key, val := t.root.min()
	return key, val, nil
}

// Max returm max element from the tree
func (t *Tree) Max() (interface{}, interface{}, error) {
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
func (n *node) add(key interface{}, value interface{}, cmp comparator) (*node, error) {
	if n == nil {
		return &node{nil, nil, key, value, 1}, nil
	}
	res, err := cmp(key, n.key)
	if err != nil {
		return nil, err
	}
	if res == 1 {
		n.right, err = n.right.add(key, value, cmp)
	} else if res == -1 {
		n.left, err = n.left.add(key, value, cmp)
	} else {
		n.key = key
	}
	return n.rebalanceTree(), err
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
func (n *node) remove(key interface{}, cmp comparator) (*node, bool, error) {
	exist := false
	if n == nil {
		return nil, exist, nil
	}
	res, err := cmp(key, n.key)
	if err != nil {
		return nil, false, err
	}
	if res == -1 {
		n.left, exist, err = n.left.remove(key, cmp)
	} else if res == 1 {
		n.right, exist, err = n.right.remove(key, cmp)
	} else {
		exist = true
		if n.left != nil && n.right != nil {
			rightMinNode := n.right.findMin()
			n.key = rightMinNode.key
			n.right, exist, err = n.right.remove(rightMinNode.key, cmp)
		} else if n.left != nil {
			n = n.left
		} else if n.right != nil {
			n = n.right
		} else {
			n = nil
			return n, exist, err
		}
	}
	return n.rebalanceTree(), exist, err
}

// get -
func (n *node) get(key interface{}, cmp comparator) (interface{}, bool, error) {
	if n == nil {
		return 0, false, nil
	}
	res, err := cmp(key, n.key)
	if err != nil {
		return nil, false, err
	}
	if res == 1 {
		return n.right.get(key, cmp)
	} else if res == -1 {
		return n.left.get(key, cmp)
	} else {
		return n.value, true, nil
	}
}

// max -
func (n *node) max() (interface{}, interface{}) {
	if n.right == nil {
		return n.key, n.value
	}
	return n.right.max()
}

// min -
func (n *node) min() (interface{}, interface{}) {
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

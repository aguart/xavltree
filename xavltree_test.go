package xavltree

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/google/btree"
)

type slice []int

type testObj struct {
	a int
	b string
}

func TestIntKeysMethods(t *testing.T) {
	var (
		err error
		// default count
		count0 int
		// count after add
		count1 int
		// count after delete
		count2 int
		// test index
		idx uint64 = 3421
		// max index
		maxIdx uint64 = 10000
		// min index
		minIdx uint64 = 0
	)
	obj := testObj{
		a: 10,
		b: "text body",
	}

	// create and fill the tree
	tree := NewTree()
	tree.TestElements(25)
	// get count
	count0 = tree.Count()

	// ==== testing add ========================================================
	err = tree.Add(idx, obj)
	if err != nil {
		t.Error("Add method:", err)
	} else {
		// get count after add
		count1 = tree.Count()
		// compare counts
		if count0 == count1 {
			t.Error("Add method:", "Error calculate count after Add")
		}
	}

	// ==== testing get ========================================================

	item, ok := tree.Get(idx)
	if !ok {
		t.Error("Get method:", "element not fount")
	} else {
		o, ok := item.(testObj)
		if !ok {
			t.Error("Get method:", "problem with type conversion")
		}
		if o != obj {
			t.Error("Get method:", "bad object")
		}
	}

	// ==== testing remove  ====================================================
	err, ok = tree.Remove(idx)
	if !ok {
		t.Error("Remove method:", "element not fount")
	} else {
		count2 = tree.Count()
		if count2 == count1 {
			t.Error("Error calculate count after Remove")
		}
	}

	// ==== testing Max method =================================================
	tree.Add(maxIdx, obj)
	key, val, err := tree.Max()
	if err != nil || key != maxIdx {
		t.Error("Max method:", err)
	} else {
		o, ok := val.(testObj)
		if !ok {
			t.Error("Max method:", "problem with type conversion")
		}
		if o != obj {
			t.Error("Max method:", "bad object")
		}
	}

	// ==== testing Min method =================================================
	tree.Add(minIdx, obj)
	key, val, err = tree.Min()
	if err != nil || key != minIdx {
		t.Error("Min method:", err)
	} else {
		o, ok := val.(testObj)
		if !ok {
			t.Error("Min method:", "problem with type conversion")
		}
		if o != obj {
			t.Error("Min method:", "bad object")
		}
	}
	PrintTree(tree.root, 0)
}

// ========== BENCHMARKS =======================================================

var (
	countBench  int    = 3000000
	findElement uint64 = 7654
)

func newSlice(count int) *slice {
	s := &slice{}
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		if i == count/2 {
			s.insert(int(findElement))
		} else {
			s.insert(r.Intn(9999))
		}
	}
	return s
}

func (s *slice) insert(value int) error {
	*s = append(*s, value)
	return nil
}

func (s *slice) find(item int) (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := sort.SearchInts(*s, item)
	return (*s)[index], true
}

type Uint64 uint64

// Less returns true if int(a) < int(b).
func (a Uint64) Less(than btree.Item) bool {
	return a < than.(Uint64)
}

func BenchmarkFindGBtree(b *testing.B) {
	tr := btree.New(btree.DefaultFreeListSize)

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < countBench; i++ {
		tr.ReplaceOrInsert(Uint64(r.Intn(9999)))
	}
	tr.ReplaceOrInsert(Uint64(findElement))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = tr.Get(Uint64(findElement))
	}
}

func BenchmarkFindxAVLTree(b *testing.B) {
	tr := NewTree()
	tr.TestElements(countBench)
	tr.Add(findElement, nil)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = tr.Get(findElement)
	}
}

func BenchmarkFindBasicLib(b *testing.B) {
	s := newSlice(countBench)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = s.find(int(findElement))
	}
}

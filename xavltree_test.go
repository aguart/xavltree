package xavltree

import (
	"math/rand"
	"sort"
	"testing"
	"time"
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
		idx int = 3421
		// max index
		maxIdx int = 10000
		// min index
		minIdx int = 0
	)
	obj := testObj{
		a: 10,
		b: "text body",
	}

	// create and fill the tree
	tree := NewIntKeys()
	tree.TestElements(20)
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

	item, ok, err := tree.Get(idx)
	if err != nil {
		t.Error("Get method:", err)
	} else {
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
	}

	// ==== testing remove  ====================================================
	err, ok = tree.Remove(idx)
	if err != nil {
		t.Error("Remove method:", err)
	} else {
		if !ok {
			t.Error("Remove method:", "element not fount")
		} else {
			count2 = tree.Count()
			if count2 == count1 {
				t.Error("Error calculate count after Remove")
			}
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
	//PrintTree(tree.root, 0)
}

func TestStringKeysMethods(t *testing.T) {
	var (
		err error
		// default count
		count0 int
		// count after add
		count1 int
		// count after delete
		count2 int
		// test index
		idx string = "testkey01"
		// max index
		maxIdx string = "zzzzzzzz"
		// min index
		minIdx string = "AAAAAAAA"
	)
	obj := testObj{
		a: 10,
		b: "text body",
	}

	// create and fill the tree
	tree := NewStrKeys()
	tree.TestElements(20)
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

	item, ok, err := tree.Get(idx)
	if err != nil {
		t.Error("Get method:", err)
	} else {
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
	}

	// ==== testing remove  ====================================================
	err, ok = tree.Remove(idx)
	if err != nil {
		t.Error("Remove method:", err)
	} else {
		if !ok {
			t.Error("Remove method:", "element not fount")
		} else {
			count2 = tree.Count()
			if count2 == count1 {
				t.Error("Error calculate count after Remove")
			}
		}
	}

	// ==== testing Max method =================================================
	tree.Add(maxIdx, obj)
	key, val, err := tree.Max()
	if err != nil || key != maxIdx {
		t.Error("Max method:", err, key)
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

	//PrintTree(tree.root, 0)
}

func TestFloat64KeysMethods(t *testing.T) {
	var (
		err error
		// default count
		count0 int
		// count after add
		count1 int
		// count after delete
		count2 int
		// test index
		idx float64 = 234.234
		// max index
		maxIdx float64 = 10000.9
		// min index
		minIdx float64 = 0.0
	)
	obj := testObj{
		a: 10,
		b: "text body",
	}

	// create and fill the tree
	tree := NewFloat64Keys()
	tree.TestElements(20)
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

	item, ok, err := tree.Get(idx)
	if err != nil {
		t.Error("Get method:", err)
	} else {
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
	}

	// ==== testing remove  ====================================================
	err, ok = tree.Remove(idx)
	if err != nil {
		t.Error("Remove method:", err)
	} else {
		if !ok {
			t.Error("Remove method:", "element not fount")
		} else {
			count2 = tree.Count()
			if count2 == count1 {
				t.Error("Error calculate count after Remove")
			}
		}
	}

	// ==== testing Max method =================================================
	tree.Add(maxIdx, obj)
	key, val, err := tree.Max()
	if err != nil || key != maxIdx {
		t.Error("Max method:", err, key)
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
	//PrintTree(tree.root, 0)
}

// ========== BENCHMARKS =======================================================

var (
	countBench  int = 300000
	findElement int = 7654
)

func newSlice(count int) *slice {
	s := &slice{}
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		if i == count/2 {
			s.insert(findElement)
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

func BenchmarkFindxAVLTree(b *testing.B) {
	t := NewIntKeys()
	t.TestElements(countBench)
	t.Add(findElement, nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = t.Get(findElement)
	}
}

func BenchmarkFindBasicLib(b *testing.B) {
	s := newSlice(countBench)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.find(findElement)
	}
}

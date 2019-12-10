package xavltree

import (
	"errors"
	"fmt"
)

type comparator func(a, b interface{}) (int, error)

// stringComparator provides a fast comparison on strings
func stringComparator(a, b interface{}) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("Recovered from: %v", r))
		}
	}()
	s1 := a.(string)
	s2 := b.(string)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1, err
	}
	if diff > 0 {
		return 1, err
	}
	return 0, err
}

// intComparator provides a basic comparison on int
func intComparator(a, b interface{}) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("Recovered from: %v", r))
		}
	}()
	ia, ib := a.(int), b.(int)
	if ia > ib {
		return 1, err
	} else if ia < ib {
		return -1, err
	} else {
		return 0, err
	}
}

// int64Comparator provides a basic comparison on int64
func int64Comparator(a, b interface{}) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("Recovered from: %v", r))
		}
	}()
	ia, ib := a.(int64), b.(int64)
	if ia > ib {
		return 1, err
	} else if ia < ib {
		return -1, err
	} else {
		return 0, err
	}
}

// float64Comparator provides a basic comparison on float64
func float64Comparator(a, b interface{}) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("Recovered from: %v", r))
		}
	}()
	ia, ib := a.(float64), b.(float64)
	if ia > ib {
		return 1, err
	} else if ia < ib {
		return -1, err
	} else {
		return 0, err
	}
}

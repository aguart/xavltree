# xavltree [![Build Status](https://travis-ci.org/aguarty/xavltree.svg?branch=master)](https://travis-ci.org/aguarty/xavltree)
Just another implementation of AVL tree.    
    
GBTree - github.com/google/btree    
xAVLTree - this package   
BasicLib - go/src/sort    
    
3.000.000 elements    

```
goos: linux
goarch: amd64
BenchmarkFindGBtree-6       10085353           122 ns/op         8 B/op        1 allocs/op
BenchmarkFindxAVLTree-6     68519924          32.5 ns/op         0 B/op        0 allocs/op
BenchmarkFindBasicLib-6     19001233          59.3 ns/op         0 B/op        0 allocs/op
```

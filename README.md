## Slice Heaps for Go

I wanted something faster than the default [containers.Heap](https://pkg.go.dev/container/heap) implementation. This library implements a simple heap structure backed by an int slice that is about 4x faster with zero allocations.


```
cpu: Intel(R) Core(TM) i7-4790K CPU @ 4.00GHz
BenchmarkContainerHeap-8   	 9540620	       128.9 ns/op	      15 B/op	       1 allocs/op
BenchmarkIntHeap-8      	31438557	        38.25 ns/op	       0 B/op	       0 allocs/op
```


## What is a heap?

A heap is a binary tree which decides a node's lineage and place based on the value of the node. The whole structure can exist as a single slice which requires less memory and is more performant than storying a map/dictionary. 

The value decides it's place in the tree, and that place is expressed as the index in the slice. Below is a min-heap (lowest value is root):

```
    1
   / \
  5   3
 / \
7   9
```

Which is stored as `[]int{1,5,3,7,9}` in memory. That is 7 is a child of 5, and 5 is a child of 1. The left most child of a node is always `2*index`. So if "1" is the root element then `1*2 = 2` is the index of 5, then `2*2 = 4` is the index of 7. If 7 had a child it's index would be `2*4 = 8` and so on.

Reads and writes are `O(log n)`. This is ideal when implementing a priority queue or min/max value cache (top 100 values from a large set).

This is a great introduction video: https://www.youtube.com/watch?v=HqPJF2L5h9U
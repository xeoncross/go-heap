package testheap

import (
	"container/heap"
	"testing"

	"github.com/xeoncross/go-heap/minheap"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkContainerHeap(b *testing.B) {

	size := 1000

	ids := make([]int, size)
	for i := 0; i < size; i++ {
		ids[i] = i
	}

	h := IntHeap(ids)
	heap.Init(&h)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heap.Pop(&h)     // delete, re-tree
		heap.Push(&h, i) // append to the end of the array
	}

	if len(ids) != size {
		b.Fail()
	}
}

func BenchmarkIntHeap(b *testing.B) {

	size := 1000

	var ids []int
	for i := 0; i < size; i++ {
		minheap.PushInt(&ids, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = minheap.PopInt(&ids) // delete, re-tree
		minheap.PushInt(&ids, i) // append to the end of the array
	}

	if len(ids) != size {
		b.Fail()
	}

}

// This is no directly comparible, but losely related and worth showing
// We have a fixed-size "ranked" map which we want to keep only the "highest" values in.
// So, as new values come along we loop through the whole array to see if something lower exists we can
// remove and replace with the new value
func BenchmarkMapIteration(b *testing.B) {

	size := 1000
	ranked := make(map[int]uint8, size)
	for i := 1; i < size; i++ {
		ranked[i] = uint8(i % 8)
	}

	var changes int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := i + size
		value := uint8(i % 8)

		sK, sV := -1, value
		for k, v := range ranked {
			if v < sV {
				sK = k
				sV = v
			}
		}

		// Found something smaller? Remove it and add this
		if sK != -1 {
			changes++
			delete(ranked, sK)
			ranked[key] = value
		}

	}

	// b.Logf("%d iterations, %d changes\n", b.N, changes)
}

// TODO: verify this is correct
// func BenchmarkRostretto(b *testing.B) {

// 	size := 1000

// 	cache, err := ristretto.NewCache(&ristretto.Config{
// 		NumCounters: int64(size), // number of keys to track frequency of (10M).
// 		MaxCost:     8,           // 255
// 		BufferItems: 64,          // number of keys per Get buffer.
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := 1; i < size; i++ {
// 		cache.Set(i, uint8(i%8), int64(i%8))
// 	}

// 	var changes int
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		key := i + size
// 		value := uint8(i % 8)

// 		if cache.Set(key, value, int64(value)) {
// 			changes++
// 		}
// 	}

// 	b.Logf("%d iterations, %d changes\n", b.N, changes)
// }

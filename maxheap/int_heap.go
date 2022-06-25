package maxheap

// Based on the implementation at https://pkg.go.dev/container/heap

// This is a min heap (first in line, earliest timestamp, etc..)
// todo: max heap (heighest weight, largest order, etc...)

func PushInt(h *[]int, x int) {
	*h = append(*h, x)
	up(*h, len(*h)-1)
}

func PopInt(h *[]int) int {
	n := len(*h) - 1
	v := (*h)[0] // get top value
	// An alternative is simply to swap here (h[0], h[n] = h[n], h[0]) and mark
	// the last element as "outside" the heap so we don't act on it (heap sort)
	(*h)[0] = (*h)[n] // move bottom to top
	(*h) = (*h)[:n]   // delete last
	down(*h, 0, n)    // push new root down
	return v
}

func up(h []int, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || h[j] < h[i] {
			break
		}
		// swap
		h[i], h[j] = h[j], h[i]
		j = i
	}
}

func down(h []int, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h[j2] > h[j1] {
			j = j2 // = 2*i + 2  // right child
		}
		if h[j] <= h[i] {
			break
		}
		// Swap(i, j)
		h[i], h[j] = h[j], h[i]
		i = j
	}
	return i > i0
}

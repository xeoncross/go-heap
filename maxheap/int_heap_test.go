package maxheap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxHeap(t *testing.T) {

	var heap []int

	for _, v := range []int{3, 5, 1, 4} {
		PushInt(&heap, v)
	}

	if !reflect.DeepEqual(heap, []int{5, 4, 1, 3}) {
		t.Fatalf("invalid heap: %v\n", heap)
	}

	max := PopInt(&heap)

	if max != 5 {
		fmt.Println(max, heap)
		t.Fail()
	}

	if !reflect.DeepEqual(heap, []int{4, 3, 1}) {
		t.Fatalf("invalid heap: %v\n", heap)
	}
}

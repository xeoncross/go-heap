package minheap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinHeap(t *testing.T) {

	var heap []int

	for _, v := range []int{3, 5, 1, 4} {
		PushInt(&heap, v)
	}

	if !reflect.DeepEqual(heap, []int{1, 4, 3, 5}) {
		t.Fatalf("invalid heap: %v\n", heap)
	}

	min := PopInt(&heap)

	if min != 1 {
		t.Fail()
		fmt.Println(min, heap)
	}

	if !reflect.DeepEqual(heap, []int{3, 4, 5}) {
		t.Fatalf("invalid heap: %v\n", heap)
	}
}

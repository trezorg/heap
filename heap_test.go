package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type item struct {
	priority int
}

func (i item) Priority() int {
	return i.priority
}

func TestHeap2FactorChilds(t *testing.T) {
	h, _ := NewMinHeap(2)
	require.Equal(t, []int{1, 2}, h.children(0))
	require.Equal(t, []int{3, 4}, h.children(1))
	require.Equal(t, []int{5, 6}, h.children(2))
	require.Equal(t, []int{7, 8}, h.children(3))
}

func TestHeap3FactorChilds(t *testing.T) {
	h, _ := NewMinHeap(3)
	require.Equal(t, []int{1, 2, 3}, h.children(0))
	require.Equal(t, []int{4, 5, 6}, h.children(1))
	require.Equal(t, []int{7, 8, 9}, h.children(2))
	require.Equal(t, []int{10, 11, 12}, h.children(3))
}

func TestHeap2FactorParent(t *testing.T) {
	h, _ := NewMinHeap(2)
	require.Equal(t, 1, h.parent(3))
	require.Equal(t, 1, h.parent(4))
	require.Equal(t, 2, h.parent(5))
	require.Equal(t, 2, h.parent(6))
	require.Equal(t, 3, h.parent(7))
	require.Equal(t, 3, h.parent(8))
	require.Equal(t, 0, h.parent(1))
	require.Equal(t, 0, h.parent(2))
}

func TestHeap3FactorParent(t *testing.T) {
	h, _ := NewMinHeap(3)
	require.Equal(t, 1, h.parent(4))
	require.Equal(t, 1, h.parent(5))
	require.Equal(t, 1, h.parent(6))
	require.Equal(t, 2, h.parent(7))
	require.Equal(t, 2, h.parent(8))
	require.Equal(t, 2, h.parent(9))
	require.Equal(t, 3, h.parent(10))
	require.Equal(t, 3, h.parent(11))
	require.Equal(t, 3, h.parent(12))
	require.Equal(t, 0, h.parent(1))
	require.Equal(t, 0, h.parent(2))
	require.Equal(t, 0, h.parent(3))
	require.Equal(t, 0, h.parent(0))
}

func TestHeap2Push(t *testing.T) {
	h, _ := NewMinHeap(2)

	h.Push(item{priority: 1})
	h.Push(item{priority: 2})
	h.Push(item{priority: 3})
	h.Push(item{priority: 4})
	h.Push(item{priority: 5})

	var priorities []int

	for _, item := range h.items {
		priorities = append(priorities, item.Priority())
	}

	require.Equal(t, []int{1, 2, 3, 4, 5}, priorities)

	h, _ = NewMinHeap(2)

	h.Push(item{priority: 5})
	h.Push(item{priority: 4})
	h.Push(item{priority: 3})
	h.Push(item{priority: 2})
	h.Push(item{priority: 1})

	priorities = priorities[:0]

	for _, item := range h.items {
		priorities = append(priorities, item.Priority())
	}

	require.Equal(t, []int{1, 2, 4, 5, 3}, priorities)

}

func TestHeap3Push(t *testing.T) {
	h, _ := NewMinHeap(3)

	h.Push(item{priority: 1})
	h.Push(item{priority: 2})
	h.Push(item{priority: 3})
	h.Push(item{priority: 4})
	h.Push(item{priority: 5})
	h.Push(item{priority: 6})
	h.Push(item{priority: 7})
	h.Push(item{priority: 8})
	h.Push(item{priority: 9})
	h.Push(item{priority: 10})

	var priorities []int

	for _, item := range h.items {
		priorities = append(priorities, item.Priority())
	}

	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, priorities)

	h, _ = NewMinHeap(3)

	h.Push(item{priority: 10})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 7})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 4})
	h.Push(item{priority: 3})
	h.Push(item{priority: 2})
	h.Push(item{priority: 1})

	require.Equal(t, 1, h.Pop().Priority())
	require.Equal(t, 2, h.Pop().Priority())
	require.Equal(t, 3, h.Pop().Priority())
	require.Equal(t, 4, h.Pop().Priority())
	require.Equal(t, 5, h.Pop().Priority())
	require.Equal(t, 6, h.Pop().Priority())
	require.Equal(t, 7, h.Pop().Priority())
	require.Equal(t, 8, h.Pop().Priority())
	require.Equal(t, 9, h.Pop().Priority())
	require.Equal(t, 10, h.Pop().Priority())

	h, _ = NewMinHeap(3)

	h.Push(item{priority: 3})
	h.Push(item{priority: 2})
	h.Push(item{priority: 7})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 10})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 4})
	h.Push(item{priority: 1})

	priorities = priorities[:0]

	require.Equal(t, 1, h.Pop().Priority())
	require.Equal(t, 2, h.Pop().Priority())
	require.Equal(t, 3, h.Pop().Priority())
	require.Equal(t, 4, h.Pop().Priority())
	require.Equal(t, 5, h.Pop().Priority())
	require.Equal(t, 6, h.Pop().Priority())
	require.Equal(t, 7, h.Pop().Priority())
	require.Equal(t, 8, h.Pop().Priority())
	require.Equal(t, 9, h.Pop().Priority())
	require.Equal(t, 10, h.Pop().Priority())

}

func TestMaxHeap3Push(t *testing.T) {

	h, _ := NewMaxHeap(3)

	h.Push(item{priority: 10})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 7})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 4})
	h.Push(item{priority: 3})
	h.Push(item{priority: 2})
	h.Push(item{priority: 1})

	require.Equal(t, 10, h.Pop().Priority())
	require.Equal(t, 9, h.Pop().Priority())
	require.Equal(t, 8, h.Pop().Priority())
	require.Equal(t, 7, h.Pop().Priority())
	require.Equal(t, 6, h.Pop().Priority())
	require.Equal(t, 5, h.Pop().Priority())
	require.Equal(t, 4, h.Pop().Priority())
	require.Equal(t, 3, h.Pop().Priority())
	require.Equal(t, 2, h.Pop().Priority())
	require.Equal(t, 1, h.Pop().Priority())

	h, _ = NewMaxHeap(3)

	h.Push(item{priority: 3})
	h.Push(item{priority: 2})
	h.Push(item{priority: 7})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 10})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 4})
	h.Push(item{priority: 1})

	require.Equal(t, 10, h.Pop().Priority())
	require.Equal(t, 9, h.Pop().Priority())
	require.Equal(t, 8, h.Pop().Priority())
	require.Equal(t, 7, h.Pop().Priority())
	require.Equal(t, 6, h.Pop().Priority())
	require.Equal(t, 5, h.Pop().Priority())
	require.Equal(t, 4, h.Pop().Priority())
	require.Equal(t, 3, h.Pop().Priority())
	require.Equal(t, 2, h.Pop().Priority())
	require.Equal(t, 1, h.Pop().Priority())

}

func TestMaxHeap3Pick(t *testing.T) {

	h, _ := NewMaxHeap(3)

	h.Push(item{priority: 10})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 7})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 4})
	h.Push(item{priority: 3})
	h.Push(item{priority: 2})
	h.Push(item{priority: 1})

	require.Equal(t, 10, h.pick().Priority())

	h, _ = NewMaxHeap(3)

	h.Push(item{priority: 3})
	h.Push(item{priority: 2})
	h.Push(item{priority: 7})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 10})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 4})
	h.Push(item{priority: 1})

	require.Equal(t, 10, h.pick().Priority())
}

func TestMaxHeap5Push(t *testing.T) {

	h, _ := NewMaxHeap(5)

	h.Push(item{priority: 3})
	h.Push(item{priority: 2})
	h.Push(item{priority: 7})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 10})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 4})
	h.Push(item{priority: 1})
	h.Push(item{priority: 11})
	h.Push(item{priority: 12})
	h.Push(item{priority: 100})
	h.Push(item{priority: 4})
	h.Push(item{priority: 1})

	require.Equal(t, 100, h.pick().Priority())
	require.Equal(t, 100, h.Pop().Priority())
	require.Equal(t, 12, h.Pop().Priority())
	require.Equal(t, 11, h.Pop().Priority())
	require.Equal(t, 10, h.Pop().Priority())

}

func TestMaxHeap3Heapify(t *testing.T) {

	h, _ := NewMaxHeap(3)

	h.Heapify(item{priority: 3}, item{priority: 4}, item{priority: 10}, item{priority: -1}, item{priority: 22})

	require.Equal(t, 22, h.Pop().Priority())
	require.Equal(t, 10, h.Pop().Priority())
	require.Equal(t, 4, h.Pop().Priority())
	require.Equal(t, 3, h.Pop().Priority())
	require.Equal(t, -1, h.Pop().Priority())

}

func TestMinHeap3Heapify(t *testing.T) {

	h, _ := NewMinHeap(3)

	h.Heapify(item{priority: 3}, item{priority: 4}, item{priority: 10}, item{priority: -1}, item{priority: 22})

	require.Equal(t, 5, h.Size())

	require.Equal(t, -1, h.Pop().Priority())
	require.Equal(t, 3, h.Pop().Priority())
	require.Equal(t, 4, h.Pop().Priority())
	require.Equal(t, 10, h.Pop().Priority())
	require.Equal(t, 22, h.Pop().Priority())

}

func TestMinHeap3GetFromEmpty(t *testing.T) {
	h, _ := NewMinHeap(3)
	require.Equal(t, nil, h.Pop())
}

func TestMaxPQ(t *testing.T) {
	h, _ := NewMaxPQ(5)
	h.Push(item{priority: 4})
	h.Push(item{priority: 2})
	h.Push(item{priority: 3})
	h.Push(item{priority: 1})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 7})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 10})

	require.Equal(t, 6, h.Pop().Priority())
	require.Equal(t, 7, h.Pop().Priority())
	require.Equal(t, 8, h.Pop().Priority())
	require.Equal(t, 9, h.Pop().Priority())
	require.Equal(t, 10, h.Pop().Priority())
	require.Equal(t, nil, h.Pop())

}

func TestMinPQ(t *testing.T) {
	h, _ := NewMinPQ(5)
	h.Push(item{priority: 4})
	h.Push(item{priority: 2})
	h.Push(item{priority: 3})
	h.Push(item{priority: 1})
	h.Push(item{priority: 6})
	h.Push(item{priority: 5})
	h.Push(item{priority: 7})
	h.Push(item{priority: 9})
	h.Push(item{priority: 8})
	h.Push(item{priority: 10})

	require.Equal(t, 5, h.Pop().Priority())
	require.Equal(t, 4, h.Pop().Priority())
	require.Equal(t, 3, h.Pop().Priority())
	require.Equal(t, 2, h.Pop().Priority())
	require.Equal(t, 1, h.Pop().Priority())
	require.Equal(t, nil, h.Pop())

}

func TestMinPQHeapify(t *testing.T) {
	h, _ := NewMinPQ(5)
	h.Heapify(
		item{priority: 4},
		item{priority: 2},
		item{priority: 3},
		item{priority: 1},
		item{priority: 6},
		item{priority: 5},
		item{priority: 7},
		item{priority: 9},
		item{priority: 8},
		item{priority: 10},
	)

	require.Equal(t, 5, h.Pop().Priority())
	require.Equal(t, 4, h.Pop().Priority())
	require.Equal(t, 3, h.Pop().Priority())
	require.Equal(t, 2, h.Pop().Priority())
	require.Equal(t, 1, h.Pop().Priority())
	require.Equal(t, nil, h.Pop())

}

func TestMaxPQHeapify(t *testing.T) {
	h, _ := NewMaxPQ(5)
	h.Heapify(
		item{priority: 4},
		item{priority: 2},
		item{priority: 3},
		item{priority: 1},
		item{priority: 6},
		item{priority: 5},
		item{priority: 7},
		item{priority: 9},
		item{priority: 8},
		item{priority: 10},
	)

	require.Equal(t, 6, h.Pop().Priority())
	require.Equal(t, 7, h.Pop().Priority())
	require.Equal(t, 8, h.Pop().Priority())
	require.Equal(t, 9, h.Pop().Priority())
	require.Equal(t, 10, h.Pop().Priority())
	require.Equal(t, nil, h.Pop())

}

func TestMaxPQOrderedSlice(t *testing.T) {
	h, _ := NewMaxPQ(5)
	h.Heapify(
		item{priority: 4},
		item{priority: 2},
		item{priority: 3},
		item{priority: 1},
		item{priority: 6},
		item{priority: 5},
		item{priority: 7},
		item{priority: 9},
		item{priority: 8},
		item{priority: 10},
	)

	var priorities []int
	for _, item := range h.OrderedSlice() {
		priorities = append(priorities, item.Priority())
	}

	require.Equal(t, []int{10, 9, 8, 7, 6}, priorities)

}

func TestMinPQOrderedSlice(t *testing.T) {
	h, _ := NewMinPQ(5)
	h.Heapify(
		item{priority: 4},
		item{priority: 2},
		item{priority: 3},
		item{priority: 1},
		item{priority: 6},
		item{priority: 5},
		item{priority: 7},
		item{priority: 9},
		item{priority: 8},
		item{priority: 10},
	)

	var priorities []int
	for _, item := range h.OrderedSlice() {
		priorities = append(priorities, item.Priority())
	}

	require.Equal(t, []int{1, 2, 3, 4, 5}, priorities)

}

func TestBlankMinPQOrderedSlice(t *testing.T) {
	h, _ := NewMinPQ(5)
	require.Len(t, h.OrderedSlice(), 0)
}

func TestBlankMaxPQOrderedSlice(t *testing.T) {
	h, _ := NewMaxPQ(5)
	require.Len(t, h.OrderedSlice(), 0)
}

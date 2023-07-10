package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Item int

func (item Item) Less(compareItem Item) bool {
	return item < compareItem
}

func TestHeap2FactorChildren(t *testing.T) {
	h, _ := NewMinHeap[int](2)
	require.Equal(t, []int{1, 2}, h.children(0))
	require.Equal(t, []int{3, 4}, h.children(1))
	require.Equal(t, []int{5, 6}, h.children(2))
	require.Equal(t, []int{7, 8}, h.children(3))
}

func TestHeap3FactorChildren(t *testing.T) {
	h, _ := NewMinHeap[int](3)
	require.Equal(t, []int{1, 2, 3}, h.children(0))
	require.Equal(t, []int{4, 5, 6}, h.children(1))
	require.Equal(t, []int{7, 8, 9}, h.children(2))
	require.Equal(t, []int{10, 11, 12}, h.children(3))
}

func TestHeap2FactorParent(t *testing.T) {
	h, _ := NewMinHeap[int](2)
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
	h, _ := NewMinHeap[int](3)
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
	h, _ := NewMinHeap[int](2)

	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5)

	require.Equal(t, []int{1, 2, 3, 4, 5}, h.items)

	h, _ = NewMinHeap[int](2)

	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	require.Equal(t, []int{1, 2, 4, 5, 3}, h.items)

}

func TestHeapComparable2Push(t *testing.T) {
	h, _ := NewComparableMinHeap[Item](2)

	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5)

	require.Equal(t, []Item{1, 2, 3, 4, 5}, h.items)

	h, _ = NewComparableMinHeap[Item](2)

	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	require.Equal(t, []Item{1, 2, 4, 5, 3}, h.items)

}

func TestHeap3Push(t *testing.T) {
	h, _ := NewMinHeap[int](3)

	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5)
	h.Push(6)
	h.Push(7)
	h.Push(8)
	h.Push(9)
	h.Push(10)

	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, h.items)

	h, _ = NewMinHeap[int](3)

	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	require.Equal(t, 1, h.Pop())
	require.Equal(t, 2, h.Pop())
	require.Equal(t, 3, h.Pop())
	require.Equal(t, 4, h.Pop())
	require.Equal(t, 5, h.Pop())
	require.Equal(t, 6, h.Pop())
	require.Equal(t, 7, h.Pop())
	require.Equal(t, 8, h.Pop())
	require.Equal(t, 9, h.Pop())
	require.Equal(t, 10, h.Pop())

	h, _ = NewMinHeap[int](3)

	h.Push(3)
	h.Push(2)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(1)
	h.Push(9)
	h.Push(8)
	h.Push(4)
	h.Push(10)

	require.Equal(t, 1, h.Pop())
	require.Equal(t, 2, h.Pop())
	require.Equal(t, 3, h.Pop())
	require.Equal(t, 4, h.Pop())
	require.Equal(t, 5, h.Pop())
	require.Equal(t, 6, h.Pop())
	require.Equal(t, 7, h.Pop())
	require.Equal(t, 8, h.Pop())
	require.Equal(t, 9, h.Pop())
	require.Equal(t, 10, h.Pop())

}

func TestMaxHeap3Push(t *testing.T) {

	h, _ := NewMaxHeap[int](3)

	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	require.Equal(t, 10, h.Pop())
	require.Equal(t, 9, h.Pop())
	require.Equal(t, 8, h.Pop())
	require.Equal(t, 7, h.Pop())
	require.Equal(t, 6, h.Pop())
	require.Equal(t, 5, h.Pop())
	require.Equal(t, 4, h.Pop())
	require.Equal(t, 3, h.Pop())
	require.Equal(t, 2, h.Pop())
	require.Equal(t, 1, h.Pop())

	h, _ = NewMaxHeap[int](3)

	h.Push(3)
	h.Push(2)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(4)
	h.Push(1)

	require.Equal(t, 10, h.Pop())
	require.Equal(t, 9, h.Pop())
	require.Equal(t, 8, h.Pop())
	require.Equal(t, 7, h.Pop())
	require.Equal(t, 6, h.Pop())
	require.Equal(t, 5, h.Pop())
	require.Equal(t, 4, h.Pop())
	require.Equal(t, 3, h.Pop())
	require.Equal(t, 2, h.Pop())
	require.Equal(t, 1, h.Pop())

}

func TestMaxHeap3Pick(t *testing.T) {

	h, _ := NewMaxHeap[int](3)

	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	require.Equal(t, 10, h.pick())

	h, _ = NewMaxHeap[int](3)

	h.Push(3)
	h.Push(2)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(4)
	h.Push(1)

	require.Equal(t, 10, h.pick())
}

func TestMaxHeap5Push(t *testing.T) {

	h, _ := NewMaxHeap[int](5)

	h.Push(3)
	h.Push(2)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(4)
	h.Push(1)
	h.Push(11)
	h.Push(12)
	h.Push(100)
	h.Push(4)
	h.Push(1)

	require.Equal(t, 100, h.pick())
	require.Equal(t, 100, h.Pop())
	require.Equal(t, 12, h.Pop())
	require.Equal(t, 11, h.Pop())
	require.Equal(t, 10, h.Pop())

}

func TestMaxHeap3Heapify(t *testing.T) {

	h, _ := NewMaxHeap[int](3)

	h.Heapify(
		3,
		4,
		10,
		-1,
		22,
	)

	require.Equal(t, 22, h.Pop())
	require.Equal(t, 10, h.Pop())
	require.Equal(t, 4, h.Pop())
	require.Equal(t, 3, h.Pop())
	require.Equal(t, -1, h.Pop())

}

func TestMaxHeap3HeapifyFloats(t *testing.T) {

	h, _ := NewMaxHeap[float64](3)

	h.Heapify(
		3,
		4,
		10,
		-1,
		22,
	)

	require.Equal(t, 22.0, h.Pop())
	require.Equal(t, 10.0, h.Pop())
	require.Equal(t, 4.0, h.Pop())
	require.Equal(t, 3.0, h.Pop())
	require.Equal(t, -1.0, h.Pop())

}
func TestMaxHeap3HeapifyStrings(t *testing.T) {

	h, _ := NewMaxHeap[string](3)

	h.Heapify(
		"3",
		"4",
		"5",
		"1",
		"2",
	)

	require.Equal(t, "5", h.Pop())
	require.Equal(t, "4", h.Pop())
	require.Equal(t, "3", h.Pop())
	require.Equal(t, "2", h.Pop())
	require.Equal(t, "1", h.Pop())

}

func TestMinHeap3Heapify(t *testing.T) {

	h, _ := NewMinHeap[int](3)

	h.Heapify(
		3,
		4,
		10,
		-1,
		22,
	)

	require.Equal(t, 5, h.Size())

	require.Equal(t, -1, h.Pop())
	require.Equal(t, 3, h.Pop())
	require.Equal(t, 4, h.Pop())
	require.Equal(t, 10, h.Pop())
	require.Equal(t, 22, h.Pop())

}

func TestMinComparableHeap3Heapify(t *testing.T) {

	h, _ := NewComparableMinHeap[Item](3)

	h.Heapify(
		3,
		4,
		10,
		-1,
		22,
	)

	require.Equal(t, 5, h.Size())

	require.Equal(t, -1, int(h.Pop()))
	require.Equal(t, 3, int(h.Pop()))
	require.Equal(t, 4, int(h.Pop()))
	require.Equal(t, 10, int(h.Pop()))
	require.Equal(t, 22, int(h.Pop()))

}

func TestMinHeap3GetFromEmpty(t *testing.T) {
	h, _ := NewMinHeap[int](3)
	require.Panics(t, func() { h.Pop() })
}

func TestEmptySize(t *testing.T) {
	minH, _ := NewMinHeap[int](3)
	require.True(t, minH.Empty())
	require.Equal(t, 0, minH.Size())
	maxH, _ := NewMaxHeap[int](3)
	require.True(t, maxH.Empty())
	require.Equal(t, 0, maxH.Size())
	maxP, _ := NewMaxPQ[int](3)
	require.True(t, maxP.Empty())
	require.Equal(t, 0, maxP.Size())
	minP, _ := NewMaxPQ[int](3)
	require.True(t, minP.Empty())
	require.Equal(t, 0, minP.Size())
}

func TestMaxPQ(t *testing.T) {
	h, _ := NewMaxPQ[int](5)
	h.Push(4)
	h.Push(2)
	h.Push(3)
	h.Push(1)
	h.Push(6)
	h.Push(5)
	h.Push(7)
	h.Push(9)
	h.Push(8)
	h.Push(10)

	require.Equal(t, 6, h.Pop())
	require.Equal(t, 7, h.Pop())
	require.Equal(t, 8, h.Pop())
	require.Equal(t, 9, h.Pop())
	require.Equal(t, 10, h.Pop())
	require.Panics(t, func() { h.Pop() })

}

func TestMinPQ(t *testing.T) {
	h, _ := NewMinPQ[int](5)
	h.Push(4)
	h.Push(2)
	h.Push(3)
	h.Push(1)
	h.Push(6)
	h.Push(5)
	h.Push(7)
	h.Push(9)
	h.Push(8)
	h.Push(10)

	require.Equal(t, 5, h.Pop())
	require.Equal(t, 4, h.Pop())
	require.Equal(t, 3, h.Pop())
	require.Equal(t, 2, h.Pop())
	require.Equal(t, 1, h.Pop())
	require.Panics(t, func() { h.Pop() })

}

func TestMinPQHeapify(t *testing.T) {
	h, _ := NewMinPQ[int](5)
	h.Heapify(
		4,
		2,
		3,
		1,
		6,
		5,
		7,
		9,
		8,
		10,
	)

	require.Equal(t, 5, h.Pop())
	require.Equal(t, 4, h.Pop())
	require.Equal(t, 3, h.Pop())
	require.Equal(t, 2, h.Pop())
	require.Equal(t, 1, h.Pop())
	require.Panics(t, func() { h.Pop() })

}

func TestMaxPQHeapify(t *testing.T) {
	h, _ := NewMaxPQ[int](5)
	h.Heapify(
		4,
		2,
		3,
		1,
		6,
		5,
		7,
		9,
		8,
		10,
	)

	require.Equal(t, 6, h.Pop())
	require.Equal(t, 7, h.Pop())
	require.Equal(t, 8, h.Pop())
	require.Equal(t, 9, h.Pop())
	require.Equal(t, 10, h.Pop())
	require.Panics(t, func() { h.Pop() })

}

func TestMaxPQOrderedSlice(t *testing.T) {
	h, _ := NewMaxPQ[int](5)
	h.Heapify(
		4,
		2,
		3,
		1,
		6,
		5,
		7,
		9,
		8,
		10,
	)

	require.Equal(t, []int{10, 9, 8, 7, 6}, h.OrderedSlice())

}

func TestMinPQOrderedSlice(t *testing.T) {
	h, _ := NewMinPQ[int](5)
	h.Heapify(
		4,
		2,
		3,
		1,
		6,
		5,
		7,
		9,
		8,
		10,
	)

	require.Equal(t, []int{1, 2, 3, 4, 5}, h.OrderedSlice())

}

func TestMinComparablePQOrderedSlice(t *testing.T) {
	h, _ := NewMinComparablePQ[Item](5)
	h.Heapify(
		4,
		2,
		3,
		1,
		6,
		5,
		7,
		9,
		8,
		10,
	)

	require.Equal(t, []Item{1, 2, 3, 4, 5}, h.OrderedSlice())

}

func TestMaxComparablePQOrderedSlice(t *testing.T) {
	h, _ := NewMaxComparablePQ[Item](5)
	h.Heapify(
		4,
		2,
		3,
		1,
		6,
		5,
		7,
		9,
		8,
		10,
	)

	require.Equal(t, []Item{10, 9, 8, 7, 6}, h.OrderedSlice())

}

func TestBlankMinPQOrderedSlice(t *testing.T) {
	h, _ := NewMinPQ[int](5)
	require.Len(t, h.OrderedSlice(), 0)
}

func TestBlankMaxPQOrderedSlice(t *testing.T) {
	h, _ := NewMaxPQ[int](5)
	require.Len(t, h.OrderedSlice(), 0)
}

func TestMaxPQ3BaseTypePush(t *testing.T) {

	type value int

	h, _ := NewMaxPQ[value](5)

	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	orderedSlice := h.OrderedSlice()

	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	slice := h.Slice()

	require.Equal(t, 5, len(orderedSlice))
	require.Equal(t, 5, len(slice))
	require.Equal(t, []value{10, 9, 8, 7, 6}, orderedSlice)
	require.Equal(t, []value{6, 7, 8, 9, 10}, slice)
}

func TestMaxHeap3BaseTypePush(t *testing.T) {

	type value int

	h, _ := NewMaxHeap[value](2)

	h.Push(10)
	h.Push(9)
	h.Push(8)
	h.Push(7)
	h.Push(6)
	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	require.Equal(t, 10, h.Size())
	slice := h.Slice()
	require.Equal(t, []value{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, slice)
}

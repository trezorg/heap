package comparable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Item int

func (item Item) Less(compareItem Item) bool {
	return item < compareItem
}

func TestHeapComparator2Push(t *testing.T) {
	h, _ := NewMinHeap[Item](2)

	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5)

	require.Equal(t, []Item{1, 2, 3, 4, 5}, h.items)

	h, _ = NewMinHeap[Item](2)

	h.Push(5)
	h.Push(4)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	require.Equal(t, []Item{1, 2, 4, 5, 3}, h.items)
}

func TestMinComparatorHeap3Heapify(t *testing.T) {
	h, _ := NewMinHeap[Item](3)

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

func TestMinComparatorPQOrderedSlice(t *testing.T) {
	h, _ := NewMinPQ[Item](5)
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

func TestMaxComparatorPQOrderedSlice(t *testing.T) {
	h, _ := NewMaxPQ[Item](5)
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

package ordered

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Comparator[T any] interface {
	Less(T) bool
}

// baseHeap heap structure with values based on Ordered generic constraint
type baseHeap[T constraints.Ordered] struct {
	check    func(item1 T, item2 T) bool
	getChild func(items []T, idx []int) int
	items    []T
	factor   int
}

// MinHeap is heap that returns element with min priority
type MinHeap[T constraints.Ordered] struct {
	baseHeap[T]
}

// MaxHeap is heap that returns element with max priority
type MaxHeap[T constraints.Ordered] struct {
	baseHeap[T]
}

// MaxPQ is maximum bounded priority queue based on Ordered generic constraint
type MaxPQ[T constraints.Ordered] struct {
	baseHeap[T]
	size int
}

// MinPQ is minimum bounded priority queue
type MinPQ[T constraints.Ordered] struct {
	baseHeap[T]
	size int
}

func newHeap[T constraints.Ordered](
	factor int,
	check func(item1 T, item2 T) bool,
	getChild func(items []T, idx []int) int,
) (baseHeap[T], error) {
	if factor < 2 {
		return baseHeap[T]{}, fmt.Errorf("wrong value for factor: %d. Cannot be less than 2", factor)
	}

	return baseHeap[T]{
		items:    make([]T, 0),
		factor:   factor,
		check:    check,
		getChild: getChild,
	}, nil
}

// NewMinHeap heap constructor
func NewMinHeap[T constraints.Ordered](factor int) (MinHeap[T], error) {
	baseHeap, err := newHeap(factor, minCheck[T], getMinChild[T])
	if err != nil {
		return MinHeap[T]{}, err
	}
	return MinHeap[T]{baseHeap}, nil
}

// NewMaxHeap heap constructor
func NewMaxHeap[T constraints.Ordered](factor int) (MaxHeap[T], error) {
	baseHeap, err := newHeap(factor, maxCheck[T], getMaxChild[T])
	if err != nil {
		return MaxHeap[T]{}, err
	}
	return MaxHeap[T]{baseHeap}, nil
}

// NewMaxPQ creates maximum priority Queue with heap factor 2
func NewMaxPQ[T constraints.Ordered](size int) (MaxPQ[T], error) {
	baseHeap, err := newHeap(2, minCheck[T], getMinChild[T])
	if err != nil {
		return MaxPQ[T]{}, err
	}
	return MaxPQ[T]{baseHeap: baseHeap, size: size}, nil
}

// NewMinPQ creates minimum priority Queue with heap factor 2
func NewMinPQ[T constraints.Ordered](size int) (MinPQ[T], error) {
	baseHeap, err := newHeap(2, maxCheck[T], getMaxChild[T])
	if err != nil {
		return MinPQ[T]{}, err
	}
	return MinPQ[T]{baseHeap: baseHeap, size: size}, nil
}

func minCheck[T constraints.Ordered](item1 T, item2 T) bool {
	return item1 < item2
}

func maxCheck[T constraints.Ordered](item1 T, item2 T) bool {
	return item1 > item2
}

func checkMinMaxIndex[T constraints.Ordered](items []T, indexes []int, check func(item1 T, item2 T) bool) int {
	if len(items) == 0 || len(indexes) == 0 {
		return -1
	}
	out := indexes[0]
	if out >= len(items) {
		return -1
	}
	res := items[out]
	for i := 1; i < len(indexes); i++ {
		idx := indexes[i]
		if idx >= len(items) {
			continue
		}
		item := items[idx]
		if check(item, res) {
			res = item
			out = idx
		}
	}
	return out
}

func getMinChild[T constraints.Ordered](items []T, idx []int) int {
	return checkMinMaxIndex(items, idx, minCheck[T])
}

func getMaxChild[T constraints.Ordered](items []T, idx []int) int {
	return checkMinMaxIndex(items, idx, maxCheck[T])
}

func (h *baseHeap[T]) children(idx int) []int {
	var res []int
	for i := 0; i < h.factor; i++ {
		child := (idx * h.factor) + i + 1
		res = append(res, child)
	}
	return res
}

func parent(idx, factor int) int {
	rest, div := idx%factor, idx/factor
	if rest == 0 {
		div--
	}
	if div < 0 {
		div = 0
	}
	return div
}

func (h *baseHeap[T]) parent(idx int) int {
	return parent(idx, h.factor)
}

func (h *baseHeap[T]) up(idx int) {
	item := h.items[idx]
	for idx >= 0 {
		parent := h.parent(idx)
		parentT := h.items[parent]
		if parent != idx && h.check(item, parentT) {
			h.items[idx] = parentT
			idx = parent
		} else {
			h.items[idx] = item
			break
		}
	}
}

func (h *baseHeap[T]) push(item T) {
	h.items = append(h.items, item)
	h.up(len(h.items) - 1)
}

func (h *baseHeap[T]) heapify(items ...T) {
	h.items = items
	firstParent := (len(items) - 1) / h.factor
	for i := firstParent; i >= 0; i-- {
		h.down(i)
	}
}

func (h *baseHeap[T]) pick() T {
	if h.empty() {
		panic("empty base heap")
	}
	return h.items[0]
}

func (h *baseHeap[T]) empty() bool {
	return len(h.items) == 0
}

func (h *baseHeap[T]) len() int {
	return len(h.items)
}

func (h *baseHeap[T]) pop() T {
	if h.empty() {
		panic("empty base heap")
	}
	item := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.down(0)
	return item
}

func (h *baseHeap[T]) down(idx int) {
	if len(h.items) == 0 {
		return
	}
	item := h.items[idx]
	for idx < len(h.items) {
		child := h.getChild(h.items, h.children(idx))
		if child == -1 {
			h.items[idx] = item
			break
		}
		childT := h.items[child]
		if child != idx && h.check(childT, item) {
			h.items[idx] = childT
			idx = child
		} else {
			h.items[idx] = item
			break
		}
	}
}

// Push adds item into heap
func (h *MinHeap[T]) Push(item T) {
	h.push(item)
}

// Push adds item into heap
func (h *MaxHeap[T]) Push(item T) {
	h.push(item)
}

// Push adds item into priority queue
func (h *MinPQ[T]) Push(item T) {
	if h.len() < h.size {
		h.push(item)
		return
	}
	if h.pick() < item {
		return
	}
	h.items[0] = item
	h.down(0)
}

// Push adds item into priority queue
func (h *MaxPQ[T]) Push(item T) {
	if h.len() < h.size {
		h.push(item)
		return
	}
	if h.pick() > item {
		return
	}
	h.items[0] = item
	h.down(0)
}

// Pop returns and deletes min value
func (h *MinHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes min value
func (h *MinPQ[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxPQ[T]) Pop() T {
	return h.pop()
}

// Pick returns min value
func (h *MinHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns min value
func (h *MinPQ[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxPQ[T]) Pick() T {
	return h.pick()
}

// Empty either heap is blank
func (h *MinHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either heap is blank
func (h *MaxHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MinPQ[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MaxPQ[T]) Empty() bool {
	return h.empty()
}

// Heapify initializes heap
func (h *MinHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify initializes  heap
func (h *MaxHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify initializes  priority queue
func (h *MinPQ[T]) Heapify(items ...T) {
	h.heapify(items[:min(h.size, len(items))]...)
	for i := h.size; i < len(items); i++ {
		if h.pick() < items[i] {
			continue
		}
		h.items[0] = items[i]
		h.down(0)
	}
}

// Heapify initializes  priority queue
func (h *MaxPQ[T]) Heapify(items ...T) {
	h.heapify(items[:min(h.size, len(items))]...)
	for i := h.size; i < len(items); i++ {
		if h.pick() > items[i] {
			continue
		}
		h.items[0] = items[i]
		h.down(0)
	}
}

// Size returns heap size
func (h *MinHeap[T]) Size() int {
	return h.len()
}

// Size returns heap size
func (h *MaxHeap[T]) Size() int {
	return h.len()
}

// Slice returns heap slice
func (h *MinHeap[T]) Slice() []T {
	return h.slice()
}

// Slice returns heap slice
func (h *MaxHeap[T]) Slice() []T {
	return h.slice()
}

// Size returns priority queue size
func (h *MinPQ[T]) Size() int {
	return h.len()
}

// Size returns priority queue size
func (h *MaxPQ[T]) Size() int {
	return h.len()
}

// Slice return slice from base heap
func (h *baseHeap[T]) slice() []T {
	res := make([]T, 0, h.len())
	for !h.empty() {
		res = append(res, h.pop())
	}
	return res
}

// Slice return slice from PQ
func (h *MinPQ[T]) Slice() []T {
	return h.slice()
}

// Slice return slice from PQ
func (h *MaxPQ[T]) Slice() []T {
	return h.slice()
}

// orderedSlicePQ return ordered slice from base heap
func (h *baseHeap[T]) orderedSlicePQ() []T {
	if h.empty() {
		return make([]T, 0)
	}
	res := make([]T, h.len())
	for i := h.len() - 1; i >= 0; i-- {
		item := h.pop()
		res[i] = item
	}
	return res
}

// OrderedSlice return ordered slice from PQ
func (h *MinPQ[T]) OrderedSlice() []T {
	return h.orderedSlicePQ()
}

// OrderedSlice return ordered slice from PQ
func (h *MaxPQ[T]) OrderedSlice() []T {
	return h.orderedSlicePQ()
}

package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Comparator[T any] interface {
	Less(T) bool
}

// baseOrderedHeap heap structure with values based on Ordered generic constraint
type baseOrderedHeap[T constraints.Ordered] struct {
	items    []T
	factor   int
	check    func(item1 T, item2 T) bool
	getChild func(items []T, idx []int) int
}

// baseComparatorHeap heap structure with values based on Comparator interface
type baseComparatorHeap[T Comparator[T]] struct {
	items    []T
	factor   int
	check    func(item1 T, item2 T) bool
	getChild func(items []T, idx []int) int
}

// MinOrderedHeap is heap that returns element with min priority
type MinOrderedHeap[T constraints.Ordered] struct {
	baseOrderedHeap[T]
}

// MinComparatorHeap is heap that returns element with min priority
type MinComparatorHeap[T Comparator[T]] struct {
	baseComparatorHeap[T]
}

// MaxHeap is heap that returns element with max priority
type MaxHeap[T constraints.Ordered] struct {
	baseOrderedHeap[T]
}

// MaxComparatorHeap is heap that returns element with max priority
type MaxComparatorHeap[T Comparator[T]] struct {
	baseComparatorHeap[T]
}

// MaxOrderedPQ is maximum bounded priority queue based on Ordered generic constraint
type MaxOrderedPQ[T constraints.Ordered] struct {
	baseOrderedHeap[T]
	size int
}

// MaxComparatorPQ is maximum bounded priority queue based on Comparator interface
type MaxComparatorPQ[T Comparator[T]] struct {
	baseComparatorHeap[T]
	size int
}

// MinOrderedPQ is minimum bounded priority queue
type MinOrderedPQ[T constraints.Ordered] struct {
	baseOrderedHeap[T]
	size int
}

// MinComparatorPQ is minimum bounded priority queue
type MinComparatorPQ[T Comparator[T]] struct {
	baseComparatorHeap[T]
	size int
}

func newOrderedHeap[T constraints.Ordered](
	factor int,
	check func(item1 T, item2 T) bool,
	getChild func(items []T, idx []int) int,
) (baseOrderedHeap[T], error) {
	if factor < 2 {
		return baseOrderedHeap[T]{}, fmt.Errorf("wrong value for factor: %d. Cannot be less than 2", factor)
	}

	return baseOrderedHeap[T]{
		items:    make([]T, 0),
		factor:   factor,
		check:    check,
		getChild: getChild,
	}, nil
}

func newComparatorHeap[T Comparator[T]](
	factor int,
	check func(item1 T, item2 T) bool,
	getChild func(items []T, idx []int) int,
) (baseComparatorHeap[T], error) {
	if factor < 2 {
		return baseComparatorHeap[T]{}, fmt.Errorf("wrong value for factor: %d. Cannot be less than 2", factor)
	}

	return baseComparatorHeap[T]{
		items:    make([]T, 0),
		factor:   factor,
		check:    check,
		getChild: getChild,
	}, nil
}

// NewOrderedMinHeap heap constructor
func NewOrderedMinHeap[T constraints.Ordered](factor int) (MinOrderedHeap[T], error) {
	baseHeap, err := newOrderedHeap(factor, minCheck[T], getMinChild[T])
	if err != nil {
		return MinOrderedHeap[T]{}, err
	}
	return MinOrderedHeap[T]{baseHeap}, nil
}

// NewComparatorMinHeap heap constructor
func NewComparatorMinHeap[T Comparator[T]](factor int) (MinComparatorHeap[T], error) {
	baseHeap, err := newComparatorHeap(factor, minComparatorCheck[T], getComparatorMinChild[T])
	if err != nil {
		return MinComparatorHeap[T]{}, err
	}
	return MinComparatorHeap[T]{baseHeap}, nil
}

// NewOrderedMaxHeap heap constructor
func NewOrderedMaxHeap[T constraints.Ordered](factor int) (MaxHeap[T], error) {
	baseHeap, err := newOrderedHeap(factor, maxCheck[T], getMaxChild[T])
	if err != nil {
		return MaxHeap[T]{}, err
	}
	return MaxHeap[T]{baseHeap}, nil
}

// NewComparatorMaxHeap heap constructor
func NewComparatorMaxHeap[T Comparator[T]](factor int) (MaxComparatorHeap[T], error) {
	baseHeap, err := newComparatorHeap(factor, maxComparatorCheck[T], getComparatorMaxChild[T])
	if err != nil {
		return MaxComparatorHeap[T]{}, err
	}
	return MaxComparatorHeap[T]{baseHeap}, nil
}

// NewMaxPQ creates maximum priority Queue
func NewMaxPQ[T constraints.Ordered](size int) (MaxOrderedPQ[T], error) {
	baseHeap, err := newOrderedHeap(size, minCheck[T], getMinChild[T])
	if err != nil {
		return MaxOrderedPQ[T]{}, err
	}
	return MaxOrderedPQ[T]{baseOrderedHeap: baseHeap, size: size}, nil
}

// NewMaxComparatorPQ creates maximum priority Queue
func NewMaxComparatorPQ[T Comparator[T]](size int) (MaxComparatorPQ[T], error) {
	baseHeap, err := newComparatorHeap(size, minComparatorCheck[T], getComparatorMinChild[T])
	if err != nil {
		return MaxComparatorPQ[T]{}, err
	}
	return MaxComparatorPQ[T]{baseComparatorHeap: baseHeap, size: size}, nil
}

// NewOrderedMinPQ creates minimum priority Queue
func NewOrderedMinPQ[T constraints.Ordered](size int) (MinOrderedPQ[T], error) {
	baseHeap, err := newOrderedHeap(size, maxCheck[T], getMaxChild[T])
	if err != nil {
		return MinOrderedPQ[T]{}, err
	}
	return MinOrderedPQ[T]{baseOrderedHeap: baseHeap, size: size}, nil
}

// NewComparatorMinPQ creates maximum priority Queue
func NewMinComparatorPQ[T Comparator[T]](size int) (MinComparatorPQ[T], error) {
	baseHeap, err := newComparatorHeap(size, maxComparatorCheck[T], getComparatorMaxChild[T])
	if err != nil {
		return MinComparatorPQ[T]{}, err
	}
	return MinComparatorPQ[T]{baseComparatorHeap: baseHeap, size: size}, nil
}

func minCheck[T constraints.Ordered](item1 T, item2 T) bool {
	return item1 < item2
}

func minComparatorCheck[T Comparator[T]](item1 T, item2 T) bool {
	return item1.Less(item2)
}

func maxCheck[T constraints.Ordered](item1 T, item2 T) bool {
	return item1 > item2
}

func maxComparatorCheck[T Comparator[T]](item1 T, item2 T) bool {
	return item2.Less(item1)
}

func checkOrderedMinMaxIndex[T constraints.Ordered](items []T, indexes []int, check func(item1 T, item2 T) bool) int {
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

func checkComparatorMinMaxIndex[T Comparator[T]](items []T, indexes []int, check func(item1 T, item2 T) bool) int {
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
	return checkOrderedMinMaxIndex(items, idx, minCheck[T])
}

func getComparatorMinChild[T Comparator[T]](items []T, idx []int) int {
	return checkComparatorMinMaxIndex(items, idx, minComparatorCheck[T])
}

func getMaxChild[T constraints.Ordered](items []T, idx []int) int {
	return checkOrderedMinMaxIndex(items, idx, maxCheck[T])
}

func getComparatorMaxChild[T Comparator[T]](items []T, idx []int) int {
	return checkComparatorMinMaxIndex(items, idx, maxComparatorCheck[T])
}

func (h *baseOrderedHeap[T]) children(idx int) []int {
	var res []int
	for i := 0; i < h.factor; i++ {
		child := (idx * h.factor) + i + 1
		res = append(res, child)
	}
	return res
}

func (h *baseComparatorHeap[T]) children(idx int) []int {
	var res []int
	for i := 0; i < h.factor; i++ {
		child := (idx * h.factor) + i + 1
		res = append(res, child)
	}
	return res
}

func (h *baseOrderedHeap[T]) parent(idx int) int {
	rest, div := idx%h.factor, idx/h.factor
	if rest == 0 {
		div--
	}
	if div < 0 {
		div = 0
	}
	return div
}

func (h *baseComparatorHeap[T]) parent(idx int) int {
	rest, div := idx%h.factor, idx/h.factor
	if rest == 0 {
		div--
	}
	if div < 0 {
		div = 0
	}
	return div
}

func (h *baseOrderedHeap[T]) up(idx int) {
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

func (h *baseComparatorHeap[T]) up(idx int) {
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

func (h *baseOrderedHeap[T]) push(item T) {
	h.items = append(h.items, item)
	h.up(len(h.items) - 1)
}

func (h *baseComparatorHeap[T]) push(item T) {
	h.items = append(h.items, item)
	h.up(len(h.items) - 1)
}

func (h *baseOrderedHeap[T]) heapify(items ...T) {
	h.items = items
	firstParent := (len(items) - 1) / h.factor
	for i := firstParent; i >= 0; i-- {
		h.down(i)
	}
}

func (h *baseComparatorHeap[T]) heapify(items ...T) {
	h.items = items
	firstParent := (len(items) - 1) / h.factor
	for i := firstParent; i >= 0; i-- {
		h.down(i)
	}
}

func (h *baseOrderedHeap[T]) pick() T {
	if h.empty() {
		panic("empty base heap")
	}
	return h.items[0]
}

func (h *baseComparatorHeap[T]) pick() T {
	if h.empty() {
		panic("empty base heap")
	}
	return h.items[0]
}

func (h *baseOrderedHeap[T]) empty() bool {
	return len(h.items) == 0
}

func (h *baseComparatorHeap[T]) empty() bool {
	return len(h.items) == 0
}

func (h *baseOrderedHeap[T]) len() int {
	return len(h.items)
}

func (h *baseComparatorHeap[T]) len() int {
	return len(h.items)
}

func (h *baseOrderedHeap[T]) pop() T {
	if h.empty() {
		panic("empty base heap")
	}
	item := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.down(0)
	return item
}

func (h *baseComparatorHeap[T]) pop() T {
	if h.empty() {
		panic("empty base heap")
	}
	item := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.down(0)
	return item
}

func (h *baseOrderedHeap[T]) down(idx int) {
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

func (h *baseComparatorHeap[T]) down(idx int) {
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
func (h *MinOrderedHeap[T]) Push(item T) {
	h.push(item)
}

// Push adds item into heap
func (h *MinComparatorHeap[T]) Push(item T) {
	h.push(item)
}

// Push adds item into heap
func (h *MaxHeap[T]) Push(item T) {
	h.push(item)
}

// Push adds item into heap
func (h *MaxComparatorHeap[T]) Push(item T) {
	h.push(item)
}

// Push adds item into priority queue
func (h *MinOrderedPQ[T]) Push(item T) {
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
func (h *MinComparatorPQ[T]) Push(item T) {
	if h.len() < h.size {
		h.push(item)
		return
	}
	if h.pick().Less(item) {
		return
	}
	h.items[0] = item
	h.down(0)
}

// Push adds item into priority queue
func (h *MaxOrderedPQ[T]) Push(item T) {
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

// Push adds item into priority queue
func (h *MaxComparatorPQ[T]) Push(item T) {
	if h.len() < h.size {
		h.push(item)
		return
	}
	if item.Less(h.pick()) {
		return
	}
	h.items[0] = item
	h.down(0)
}

// Pop returns and deletes min value
func (h *MinOrderedHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes min value
func (h *MinComparatorHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxComparatorHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes min value
func (h *MinOrderedPQ[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes min value
func (h *MinComparatorPQ[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxOrderedPQ[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxComparatorPQ[T]) Pop() T {
	return h.pop()
}

// Pick returns min value
func (h *MinOrderedHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns min value
func (h *MinComparatorHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxComparatorHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns min value
func (h *MinOrderedPQ[T]) Pick() T {
	return h.pick()
}

// Pick returns min value
func (h *MinComparatorPQ[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxOrderedPQ[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxComparatorPQ[T]) Pick() T {
	return h.pick()
}

// Empty either heap is blank
func (h *MinOrderedHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either heap is blank
func (h *MinComparatorHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either heap is blank
func (h *MaxHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either heap is blank
func (h *MaxComparatorHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MinOrderedPQ[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MinComparatorPQ[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MaxOrderedPQ[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MaxComparatorPQ[T]) Empty() bool {
	return h.empty()
}

// Heapify inits heap
func (h *MinOrderedHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify inits heap
func (h *MinComparatorHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify inits heap
func (h *MaxHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify inits heap
func (h *MaxComparatorHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify inits priority queue
func (h *MinOrderedPQ[T]) Heapify(items ...T) {
	h.heapify(items[:h.size]...)
	for i := h.size; i < len(items); i++ {
		if h.pick() < items[i] {
			continue
		}
		h.items[0] = items[i]
		h.down(0)
	}
}

// Heapify inits priority queue
func (h *MinComparatorPQ[T]) Heapify(items ...T) {
	h.heapify(items[:h.size]...)
	for i := h.size; i < len(items); i++ {
		if h.pick().Less(items[i]) {
			continue
		}
		h.items[0] = items[i]
		h.down(0)
	}
}

// Heapify inits priority queue
func (h *MaxOrderedPQ[T]) Heapify(items ...T) {
	h.heapify(items[:h.size]...)
	for i := h.size; i < len(items); i++ {
		if h.pick() > items[i] {
			continue
		}
		h.items[0] = items[i]
		h.down(0)
	}
}

// Heapify inits priority queue
func (h *MaxComparatorPQ[T]) Heapify(items ...T) {
	h.heapify(items[:h.size]...)
	for i := h.size; i < len(items); i++ {
		if items[i].Less(h.pick()) {
			continue
		}
		h.items[0] = items[i]
		h.down(0)
	}
}

// Size returns heap size
func (h *MinOrderedHeap[T]) Size() int {
	return h.len()
}

// Size returns heap size
func (h *MinComparatorHeap[T]) Size() int {
	return h.len()
}

// Size returns heap size
func (h *MaxHeap[T]) Size() int {
	return h.len()
}

// Size returns heap size
func (h *MaxComparatorHeap[T]) Size() int {
	return h.len()
}

// Slice returns heap slice
func (h *MinOrderedHeap[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice returns heap slice
func (h *MinComparatorHeap[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice returns heap slice
func (h *MaxHeap[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice returns heap slice
func (h *MaxComparatorHeap[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Size returns priority queue size
func (h *MinOrderedPQ[T]) Size() int {
	return h.len()
}

// Size returns priority queue size
func (h *MinComparatorPQ[T]) Size() int {
	return h.len()
}

// Size returns priority queue size
func (h *MaxOrderedPQ[T]) Size() int {
	return h.len()
}

// Size returns priority queue size
func (h *MaxComparatorPQ[T]) Size() int {
	return h.len()
}

// Slice return slice from PQ
func (h *MinOrderedPQ[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice return slice from PQ
func (h *MinComparatorPQ[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice return slice from PQ
func (h *MaxOrderedPQ[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice return slice from PQ
func (h *MaxComparatorPQ[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// OrderedSlice return ordered slice from PQ
func (h *MinOrderedPQ[T]) OrderedSlice() []T {
	if h.Empty() {
		return make([]T, 0)
	}
	res := make([]T, h.Size())
	for i := h.Size() - 1; i >= 0; i-- {
		item := h.Pop()
		res[i] = item
	}
	return res
}

// OrderedSlice return ordered slice from PQ
func (h *MinComparatorPQ[T]) OrderedSlice() []T {
	if h.Empty() {
		return make([]T, 0)
	}
	res := make([]T, h.Size())
	for i := h.Size() - 1; i >= 0; i-- {
		item := h.Pop()
		res[i] = item
	}
	return res
}

// OrderedSlice return ordered slice from PQ
func (h *MaxOrderedPQ[T]) OrderedSlice() []T {
	if h.Empty() {
		return make([]T, 0)
	}
	res := make([]T, h.Size())
	for i := h.Size() - 1; i >= 0; i-- {
		item := h.Pop()
		res[i] = item
	}
	return res
}

// OrderedSlice return ordered slice from PQ
func (h *MaxComparatorPQ[T]) OrderedSlice() []T {
	if h.Empty() {
		return make([]T, 0)
	}
	res := make([]T, h.Size())
	for i := h.Size() - 1; i >= 0; i-- {
		item := h.Pop()
		res[i] = item
	}
	return res
}

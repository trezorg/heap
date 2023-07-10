package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Comparator[T any] interface {
	Less(T) bool
}

// baseHeap structure
type baseHeap[T constraints.Ordered] struct {
	items    []T
	factor   int
	check    func(item1 T, item2 T) bool
	getChild func(items []T, idx []int) int
}

// baseComparableHeap comparable structure
type baseComparableHeap[T Comparator[T]] struct {
	items    []T
	factor   int
	check    func(item1 T, item2 T) bool
	getChild func(items []T, idx []int) int
}

// MinHeap is heap that returns element with min priority
type MinHeap[T constraints.Ordered] struct {
	baseHeap[T]
}

// MinComparableHeap is heap that returns element with min priority
type MinComparableHeap[T Comparator[T]] struct {
	baseComparableHeap[T]
}

// MaxHeap is heap that returns element with max priority
type MaxHeap[T constraints.Ordered] struct {
	baseHeap[T]
}

// MaxComparableHeap is heap that returns element with min priority
type MaxComparableHeap[T Comparator[T]] struct {
	baseComparableHeap[T]
}

// MaxPQ is maximum bounded priority queue
type MaxPQ[T constraints.Ordered] struct {
	baseHeap[T]
	size int
}

// MaxComparablePQ is maximum bounded priority queue
type MaxComparablePQ[T Comparator[T]] struct {
	baseComparableHeap[T]
	size int
}

// MinPQ is minimum bounded priority queue
type MinPQ[T constraints.Ordered] struct {
	baseHeap[T]
	size int
}

// MinComparablePQ is minimum bounded priority queue
type MinComparablePQ[T Comparator[T]] struct {
	baseComparableHeap[T]
	size int
}

func newHeap[T constraints.Ordered](factor int,
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

func newComparableHeap[T Comparator[T]](factor int,
	check func(item1 T, item2 T) bool,
	getChild func(items []T, idx []int) int,
) (baseComparableHeap[T], error) {
	if factor < 2 {
		return baseComparableHeap[T]{}, fmt.Errorf("wrong value for factor: %d. Cannot be less than 2", factor)
	}

	return baseComparableHeap[T]{
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

// NewComparableMinHeap heap constructor
func NewComparableMinHeap[T Comparator[T]](factor int) (MinComparableHeap[T], error) {
	baseHeap, err := newComparableHeap(factor, minComporableCheck[T], getComparableMinChild[T])
	if err != nil {
		return MinComparableHeap[T]{}, err
	}
	return MinComparableHeap[T]{baseHeap}, nil
}

// NewMaxHeap heap constructor
func NewMaxHeap[T constraints.Ordered](factor int) (MaxHeap[T], error) {
	baseHeap, err := newHeap(factor, maxCheck[T], getMaxChild[T])
	if err != nil {
		return MaxHeap[T]{}, err
	}
	return MaxHeap[T]{baseHeap}, nil
}

// NewComparableMaxHeap heap constructor
func NewComparableMaxHeap[T Comparator[T]](factor int) (MaxComparableHeap[T], error) {
	baseHeap, err := newComparableHeap(factor, maxComporableCheck[T], getComparableMaxChild[T])
	if err != nil {
		return MaxComparableHeap[T]{}, err
	}
	return MaxComparableHeap[T]{baseHeap}, nil
}

// NewMaxPQ creates maximum priority Queue
func NewMaxPQ[T constraints.Ordered](size int) (MaxPQ[T], error) {
	baseHeap, err := newHeap(size, minCheck[T], getMinChild[T])
	if err != nil {
		return MaxPQ[T]{}, err
	}
	return MaxPQ[T]{baseHeap: baseHeap, size: size}, nil
}

// NewMaxComparablePQ creates maximum priority Queue
func NewMaxComparablePQ[T Comparator[T]](size int) (MaxComparablePQ[T], error) {
	baseHeap, err := newComparableHeap(size, minComporableCheck[T], getComparableMinChild[T])
	if err != nil {
		return MaxComparablePQ[T]{}, err
	}
	return MaxComparablePQ[T]{baseComparableHeap: baseHeap, size: size}, nil
}

// NewMinPQ creates minimum priority Queue
func NewMinPQ[T constraints.Ordered](size int) (MinPQ[T], error) {
	baseHeap, err := newHeap(size, maxCheck[T], getMaxChild[T])
	if err != nil {
		return MinPQ[T]{}, err
	}
	return MinPQ[T]{baseHeap: baseHeap, size: size}, nil
}

// NewComparableMinPQ creates maximum priority Queue
func NewMinComparablePQ[T Comparator[T]](size int) (MinComparablePQ[T], error) {
	baseHeap, err := newComparableHeap(size, maxComporableCheck[T], getComparableMaxChild[T])
	if err != nil {
		return MinComparablePQ[T]{}, err
	}
	return MinComparablePQ[T]{baseComparableHeap: baseHeap, size: size}, nil
}

func minCheck[T constraints.Ordered](item1 T, item2 T) bool {
	return item1 < item2
}

func minComporableCheck[T Comparator[T]](item1 T, item2 T) bool {
	return item1.Less(item2)
}

func maxCheck[T constraints.Ordered](item1 T, item2 T) bool {
	return item1 > item2
}

func maxComporableCheck[T Comparator[T]](item1 T, item2 T) bool {
	return item2.Less(item1)
}

func checkMinMaxIndex[T constraints.Ordered](items []T, indexes []int, check func(item1 T, item2 T) bool) int {
	if len(items) == 0 {
		return -1
	}
	var aliveIdx []int
	for i := 0; i < len(indexes); i++ {
		if indexes[i] < len(items) {
			aliveIdx = append(aliveIdx, indexes[i])
		}
	}
	if len(aliveIdx) == 0 {
		return -1
	}

	res := items[aliveIdx[0]]
	out := aliveIdx[0]
	for i := 1; i < len(aliveIdx); i++ {
		item := items[aliveIdx[i]]
		if check(item, res) {
			res = item
			out = aliveIdx[i]
		}
	}
	return out
}

func checkComporableMinMaxIndex[T Comparator[T]](items []T, indexes []int, check func(item1 T, item2 T) bool) int {
	if len(items) == 0 {
		return -1
	}
	var aliveIdx []int
	for i := 0; i < len(indexes); i++ {
		if indexes[i] < len(items) {
			aliveIdx = append(aliveIdx, indexes[i])
		}
	}
	if len(aliveIdx) == 0 {
		return -1
	}

	res := items[aliveIdx[0]]
	out := aliveIdx[0]
	for i := 1; i < len(aliveIdx); i++ {
		item := items[aliveIdx[i]]
		if check(item, res) {
			res = item
			out = aliveIdx[i]
		}
	}
	return out
}

func getMinChild[T constraints.Ordered](items []T, idx []int) int {
	return checkMinMaxIndex(items, idx, minCheck[T])
}

func getComparableMinChild[T Comparator[T]](items []T, idx []int) int {
	return checkComporableMinMaxIndex(items, idx, minComporableCheck[T])
}

func getMaxChild[T constraints.Ordered](items []T, idx []int) int {
	return checkMinMaxIndex(items, idx, maxCheck[T])
}

func getComparableMaxChild[T Comparator[T]](items []T, idx []int) int {
	return checkComporableMinMaxIndex(items, idx, maxComporableCheck[T])
}

func (h *baseHeap[T]) children(idx int) []int {
	var res []int
	for i := 0; i < h.factor; i++ {
		child := (idx * h.factor) + i + 1
		res = append(res, child)
	}
	return res
}

func (h *baseComparableHeap[T]) children(idx int) []int {
	var res []int
	for i := 0; i < h.factor; i++ {
		child := (idx * h.factor) + i + 1
		res = append(res, child)
	}
	return res
}

func (h *baseHeap[T]) parent(idx int) int {
	rest, div := idx%h.factor, idx/h.factor
	if rest == 0 {
		div--
	}
	if div < 0 {
		div = 0
	}
	return div
}

func (h *baseComparableHeap[T]) parent(idx int) int {
	rest, div := idx%h.factor, idx/h.factor
	if rest == 0 {
		div--
	}
	if div < 0 {
		div = 0
	}
	return div
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

func (h *baseComparableHeap[T]) up(idx int) {
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

func (h *baseComparableHeap[T]) push(item T) {
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

func (h *baseComparableHeap[T]) heapify(items ...T) {
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

func (h *baseComparableHeap[T]) pick() T {
	if h.empty() {
		panic("empty base heap")
	}
	return h.items[0]
}

func (h *baseHeap[T]) empty() bool {
	return len(h.items) == 0
}

func (h *baseComparableHeap[T]) empty() bool {
	return len(h.items) == 0
}

func (h *baseHeap[T]) len() int {
	return len(h.items)
}

func (h *baseComparableHeap[T]) len() int {
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

func (h *baseComparableHeap[T]) pop() T {
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

func (h *baseComparableHeap[T]) down(idx int) {
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
func (h *MinComparableHeap[T]) Push(item T) {
	h.push(item)
}

// Push adds item into heap
func (h *MaxHeap[T]) Push(item T) {
	h.push(item)
}

// Push adds item into heap
func (h *MaxComparableHeap[T]) Push(item T) {
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
func (h *MinComparablePQ[T]) Push(item T) {
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

// Push adds item into priority queue
func (h *MaxComparablePQ[T]) Push(item T) {
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
func (h *MinHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes min value
func (h *MinComparableHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxComparableHeap[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes min value
func (h *MinPQ[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes min value
func (h *MinComparablePQ[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxPQ[T]) Pop() T {
	return h.pop()
}

// Pop returns and deletes max value
func (h *MaxComparablePQ[T]) Pop() T {
	return h.pop()
}

// Pick returns min value
func (h *MinHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns min value
func (h *MinComparableHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxComparableHeap[T]) Pick() T {
	return h.pick()
}

// Pick returns min value
func (h *MinPQ[T]) Pick() T {
	return h.pick()
}

// Pick returns min value
func (h *MinComparablePQ[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxPQ[T]) Pick() T {
	return h.pick()
}

// Pick returns max value
func (h *MaxComparablePQ[T]) Pick() T {
	return h.pick()
}

// Empty either heap is blank
func (h *MinHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either heap is blank
func (h *MinComparableHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either heap is blank
func (h *MaxHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either heap is blank
func (h *MaxComparableHeap[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MinPQ[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MinComparablePQ[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MaxPQ[T]) Empty() bool {
	return h.empty()
}

// Empty either priority queue is blank
func (h *MaxComparablePQ[T]) Empty() bool {
	return h.empty()
}

// Heapify inits heap
func (h *MinHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify inits heap
func (h *MinComparableHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify inits heap
func (h *MaxHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify inits heap
func (h *MaxComparableHeap[T]) Heapify(items ...T) {
	h.heapify(items...)
}

// Heapify inits priority queue
func (h *MinPQ[T]) Heapify(items ...T) {
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
func (h *MinComparablePQ[T]) Heapify(items ...T) {
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
func (h *MaxPQ[T]) Heapify(items ...T) {
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
func (h *MaxComparablePQ[T]) Heapify(items ...T) {
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
func (h *MinHeap[T]) Size() int {
	return h.len()
}

// Size returns heap size
func (h *MinComparableHeap[T]) Size() int {
	return h.len()
}

// Size returns heap size
func (h *MaxHeap[T]) Size() int {
	return h.len()
}

// Size returns heap size
func (h *MaxComparableHeap[T]) Size() int {
	return h.len()
}

// Slice returns heap slice
func (h *MinHeap[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice returns heap slice
func (h *MinComparableHeap[T]) Slice() []T {
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
func (h *MaxComparableHeap[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Size returns priority queue size
func (h *MinPQ[T]) Size() int {
	return h.len()
}

// Size returns priority queue size
func (h *MinComparablePQ[T]) Size() int {
	return h.len()
}

// Size returns priority queue size
func (h *MaxPQ[T]) Size() int {
	return h.len()
}

// Size returns priority queue size
func (h *MaxComparablePQ[T]) Size() int {
	return h.len()
}

// Slice return slice from PQ
func (h *MinPQ[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice return slice from PQ
func (h *MinComparablePQ[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice return slice from PQ
func (h *MaxPQ[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// Slice return slice from PQ
func (h *MaxComparablePQ[T]) Slice() []T {
	res := make([]T, 0, h.Size())
	for !h.Empty() {
		res = append(res, h.Pop())
	}
	return res
}

// OrderedSlice return ordered slice from PQ
func (h *MinPQ[T]) OrderedSlice() []T {
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
func (h *MinComparablePQ[T]) OrderedSlice() []T {
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
func (h *MaxPQ[T]) OrderedSlice() []T {
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
func (h *MaxComparablePQ[T]) OrderedSlice() []T {
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

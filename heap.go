package heap

import "fmt"

// Item heap item
type Item interface {
	Priority() int
}

// baseHeap structure
type baseHeap struct {
	items    []Item
	factor   int
	check    func(item1 int, item2 int) bool
	getChild func(items []Item, idx []int) int
}

// MinHeap is heap that returns element with min priority
type MinHeap struct {
	baseHeap
}

// MaxHeap is heap that returns element with max priority
type MaxHeap struct {
	baseHeap
}

// MaxPQ is maximum bounded priority queue
type MaxPQ struct {
	baseHeap
	size int
}

// MinPQ is minimum bounded priority queue
type MinPQ struct {
	baseHeap
	size int
}

func newHeap(factor int,
	check func(item1 int, item2 int) bool,
	getChild func(items []Item, idx []int) int,
) (baseHeap, error) {
	if factor < 2 {
		return baseHeap{}, fmt.Errorf("wrong value for factor: %d. Cannot be less than 2", factor)
	}

	return baseHeap{
		items:    make([]Item, 0),
		factor:   factor,
		check:    check,
		getChild: getChild,
	}, nil
}

// NewMinHeap heap constructor
func NewMinHeap(factor int) (MinHeap, error) {
	baseHeap, err := newHeap(factor, minCheck, getMinChild)
	if err != nil {
		return MinHeap{}, err
	}
	return MinHeap{baseHeap}, nil
}

// NewMaxHeap heap constructor
func NewMaxHeap(factor int) (MaxHeap, error) {
	baseHeap, err := newHeap(factor, maxCheck, getMaxChild)
	if err != nil {
		return MaxHeap{}, err
	}
	return MaxHeap{baseHeap}, nil
}

// NewMaxPQ creates maximum priority Queue
func NewMaxPQ(size int) (MaxPQ, error) {
	baseHeap, err := newHeap(2, minCheck, getMinChild)
	if err != nil {
		return MaxPQ{}, err
	}
	return MaxPQ{baseHeap: baseHeap, size: size}, nil
}

// NewMinPQ creates minimum priority Queue
func NewMinPQ(size int) (MinPQ, error) {
	baseHeap, err := newHeap(2, maxCheck, getMaxChild)
	if err != nil {
		return MinPQ{}, err
	}
	return MinPQ{baseHeap: baseHeap, size: size}, nil
}

func minCheck(item1 int, item2 int) bool {
	return item1 < item2
}

func maxCheck(item1 int, item2 int) bool {
	return item1 > item2
}

func checkMinMaxIndex(items []Item, indexes []int, check func(item1 int, item2 int) bool) int {
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
		if check(item.Priority(), res.Priority()) {
			res = item
			out = aliveIdx[i]
		}
	}
	return out
}

func getMinChild(items []Item, idx []int) int {
	return checkMinMaxIndex(items, idx, minCheck)
}

func getMaxChild(items []Item, idx []int) int {
	return checkMinMaxIndex(items, idx, maxCheck)
}

func (h *baseHeap) children(idx int) []int {
	var res []int
	for i := 0; i < h.factor; i++ {
		child := (idx * h.factor) + i + 1
		res = append(res, child)
	}
	return res
}

func (h *baseHeap) parent(idx int) int {
	rest, div := idx%h.factor, idx/h.factor
	if rest == 0 {
		div--
	}
	if div < 0 {
		div = 0
	}
	return div
}

func (h *baseHeap) up(idx int) {
	item := h.items[idx]
	for idx >= 0 {
		parent := h.parent(idx)
		parentItem := h.items[parent]
		if parent != idx && h.check(item.Priority(), parentItem.Priority()) {
			h.items[idx] = parentItem
			idx = parent
		} else {
			h.items[idx] = item
			break
		}
	}
}

func (h *baseHeap) push(item Item) {
	h.items = append(h.items, item)
	h.up(len(h.items) - 1)
}

func (h *baseHeap) heapify(items ...Item) {
	h.items = items
	firstParent := (len(items) - 1) / h.factor
	for i := firstParent; i >= 0; i-- {
		h.down(i)
	}
}

func (h *baseHeap) pick() Item {
	if h.empty() {
		return nil
	}
	return h.items[0]
}

func (h *baseHeap) empty() bool {
	return len(h.items) == 0
}

func (h *baseHeap) len() int {
	return len(h.items)
}

func (h *baseHeap) pop() Item {
	if h.empty() {
		return nil
	}
	item := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.down(0)
	return item
}

func (h *baseHeap) down(idx int) {
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
		childItem := h.items[child]
		if child != idx && h.check(childItem.Priority(), item.Priority()) {
			h.items[idx] = childItem
			idx = child
		} else {
			h.items[idx] = item
			break
		}
	}
}

// Push adds item into heap
func (h *MinHeap) Push(item Item) {
	h.baseHeap.push(item)
}

// Push adds item into heap
func (h *MaxHeap) Push(item Item) {
	h.baseHeap.push(item)
}

// Push adds item into priority queue
func (h *MinPQ) Push(item Item) {
	if h.baseHeap.len() < h.size {
		h.baseHeap.push(item)
		return
	}
	if h.baseHeap.pick().Priority() < item.Priority() {
		return
	}
	h.items[0] = item
	h.down(0)
}

// Push adds item into priority queue
func (h *MaxPQ) Push(item Item) {
	if h.baseHeap.len() < h.size {
		h.baseHeap.push(item)
		return
	}
	if h.baseHeap.pick().Priority() > item.Priority() {
		return
	}
	h.items[0] = item
	h.down(0)
}

// Pop returns and deletes min value
func (h *MinHeap) Pop() Item {
	return h.baseHeap.pop()
}

// Pop returns and deletes max value
func (h *MaxHeap) Pop() Item {
	return h.baseHeap.pop()
}

// Pop returns and deletes max value
func (h *MaxPQ) Pop() Item {
	return h.baseHeap.pop()
}

// Pop returns and deletes max value
func (h *MinPQ) Pop() Item {
	return h.baseHeap.pop()
}

// Pick returns min value
func (h *MinHeap) Pick() Item {
	return h.baseHeap.pick()
}

// Pick returns max value
func (h *MaxHeap) Pick() Item {
	return h.baseHeap.pick()
}

// Pick returns max value
func (h *MaxPQ) Pick() Item {
	return h.baseHeap.pick()
}

// Pick returns min value
func (h *MinPQ) Pick() Item {
	return h.baseHeap.pick()
}

// Empty either heap is blank
func (h *MaxHeap) Empty() bool {
	return h.baseHeap.empty()
}

// Empty either heap is blank
func (h *MinHeap) Empty() bool {
	return h.baseHeap.empty()
}

// Empty either priority queue is blank
func (h *MinPQ) Empty() bool {
	return h.baseHeap.empty()
}

// Empty either priority queue is blank
func (h *MaxPQ) Empty() bool {
	return h.baseHeap.empty()
}

// Heapify inits heap
func (h *MaxHeap) Heapify(items ...Item) {
	h.baseHeap.heapify(items...)
}

// Heapify inits heap
func (h *MinHeap) Heapify(items ...Item) {
	h.baseHeap.heapify(items...)
}

// Heapify inits priority queue
func (h *MinPQ) Heapify(items ...Item) {
	h.baseHeap.heapify(items[:h.size]...)
	for i := h.size; i < len(items); i++ {
		if h.baseHeap.pick().Priority() < items[i].Priority() {
			return
		}
		h.items[0] = items[i]
		h.down(0)
	}
}

// Heapify inits priority queue
func (h *MaxPQ) Heapify(items ...Item) {
	h.baseHeap.heapify(items[:h.size]...)
	for i := h.size; i < len(items); i++ {
		if h.baseHeap.pick().Priority() > items[i].Priority() {
			return
		}
		h.items[0] = items[i]
		h.down(0)
	}
}

// Size returns heap size
func (h *MaxHeap) Size() int {
	return h.baseHeap.len()
}

// Size returns heap size
func (h *MinHeap) Size() int {
	return h.baseHeap.len()
}

// Size returns priority queue size
func (h *MaxPQ) Size() int {
	return h.baseHeap.len()
}

// Size returns priority queue size
func (h *MinPQ) Size() int {
	return h.baseHeap.len()
}

// OrderedSlice return ordered slice from PQ
func (h *MaxPQ) OrderedSlice() []Item {
	res := make([]Item, h.Size(), h.Size())
	for i, item := h.Size()-1, h.Pop(); i >= 0; item, i = h.Pop(), i-1 {
		res[i] = item
	}
	return res
}

// OrderedSlice return ordered slice from PQ
func (h *MinPQ) OrderedSlice() []Item {
	res := make([]Item, h.Size(), h.Size())
	for i, item := h.Size()-1, h.Pop(); i >= 0; item, i = h.Pop(), i-1 {
		res[i] = item
	}
	return res
}

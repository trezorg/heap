package examples

import "github.com/trezorg/heap/ordered"

func maxheap() {
	h, _ := ordered.NewMaxHeap[string](3)

	h.Push("a")
	h.Push("b")
	h.Push("c")

	println(h.Pop()) // "c"
	println(h.Pop()) // "b"
	println(h.Pop()) // "a"
}

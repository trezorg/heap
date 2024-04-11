package main

import "github.com/trezorg/heap/ordered"

func main() {
	h, _ := ordered.NewMinHeap[string](3)

	h.Push("c")
	h.Push("b")
	h.Push("a")

	println(h.Pop()) // "a"
	println(h.Pop()) // "b"
	println(h.Pop()) // "c"
}

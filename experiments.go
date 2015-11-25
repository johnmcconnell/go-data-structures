package main

import (
	"container/heap"
	"container/list"
)

const (
	// Byte ...
	Byte = 1
	// KiloByte ...
	KiloByte = 1024 * Byte
	// MegaByte ...
	MegaByte = 1024 * KiloByte
	// SmallSize ...
	SmallSize = 100 * Byte
	// BigSize ...
	BigSize = 100 * MegaByte
)

func main() {
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

// Len ...
func (h IntHeap) Len() int {
	return len(h)
}

// Less ...
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap ...
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push ...
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop ...
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// An StrHeap is a min-heap of ints.
type StrHeap []string

// Len ...
func (h StrHeap) Len() int {
	return len(h)
}

// Less ...
func (h StrHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap ...
func (h StrHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push ...
func (h *StrHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

// Pop ...
func (h *StrHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// IntHeapN ...
func IntHeapN(Size int) *IntHeap {
	H := &IntHeap{}
	heap.Init(H)

	for i := 0; i < Size; i++ {
		heap.Push(H, i+1)
	}

	return H
}

// SmallIntHeap ...
func SmallIntHeap() *IntHeap {
	return IntHeapN(SmallSize)
}

// BigIntHeap ...
func BigIntHeap() *IntHeap {
	return IntHeapN(BigSize)
}

// Slice ...
func Slice(Size int) []int {
	Elements := make([]int, Size)

	for i := range Elements {
		Elements[i] = i + 1
	}

	return Elements
}

// SmallSlice ...
func SmallSlice() []int {
	return Slice(SmallSize)
}

// BigSlice ...
func BigSlice() []int {
	return Slice(BigSize)
}

// SmallSliceP ...
func SmallSliceP() *[]int {
	Elements := SmallSlice()

	return &Elements
}

// LinkedList ...
func LinkedList(Size int) *list.List {
	List := list.New()

	for i := 0; i < Size; i++ {
		List.PushBack(i + 1)
	}

	return List
}

// SmallLinkedList ...
func SmallLinkedList() *list.List {
	return LinkedList(SmallSize)
}

// BigLinkedList ...
func BigLinkedList() *list.List {
	return LinkedList(BigSize)
}

// Int ...
type Int int

// S ...
type S struct {
	V int
	B int
}

// Addable ...
type Addable interface {
	Add(Addable) Addable
}

// Add ...
func (i Int) Add(j Addable) Addable {
	return i + j.(Int)
}

// Add ...
func (i S) Add(j Addable) Addable {
	s := j.(S)

	v := S{
		V: i.V + s.V,
	}

	return v
}

// AddI ...
func AddI(x Int, y Int) Int {
	return x + y
}

// AddS ...
func AddS(x S, y S) S {
	return S{x.V + y.V, 0}
}

// Add ...
func Add(x int, y int) int {
	return x + y
}

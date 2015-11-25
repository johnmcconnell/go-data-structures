package main

import (
	"container/heap"
	"testing"
)

func BenchmarkEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkCPU(b *testing.B) {
	X := 43

	for i := 0; i < b.N; i++ {
		if X%2 == 0 {
			X = X*3 + 1
		} else {
			X = X / 2
		}
	}
}

func BenchmarkSmallSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List := SmallSlice()

		s := 0

		for a := range List {
			s += List[a]
		}
	}
}

func BenchmarkSmallSliceP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ListP := SmallSliceP()
		List := *ListP

		s := 0

		for a := range List {
			s += List[a]
		}
	}
}

func BenchmarkSmallLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		List := SmallLinkedList()

		s := 0

		e := List.Front()
		for e != nil {
			s += e.Value.(int)

			e = e.Next()
		}
	}
}

func BenchmarkSmallSliceRead(b *testing.B) {
	List := SmallSlice()

	s := 0
	a := 0
	l := len(List)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = a % l
		s += List[a]
	}
}

func BenchmarkBigSliceRead(b *testing.B) {
	List := BigSlice()

	s := 0
	a := 0
	l := len(List)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = a % l
		s += List[a]
	}
}

func BenchmarkSmallLinkedListRead(b *testing.B) {
	List := SmallLinkedList()

	e := List.Front()
	s := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if e == nil {
			e = List.Front()
		}

		s += e.Value.(int)
		e = e.Next()
	}
}

func BenchmarkBigLinkedListRead(b *testing.B) {
	List := BigLinkedList()

	e := List.Front()
	s := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if e == nil {
			e = List.Front()
		}

		s += e.Value.(int)
		e = e.Next()
	}
}

func BenchmarkInterfaceExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var A Int = 4
		var B Int = 7

		A.Add(B)
	}
}

func BenchmarkStructInterfaceExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		A := S{4, 0}
		B := S{7, 0}

		A.Add(B)
	}
}

func BenchmarkSimpleAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X := 4
		Y := 7

		Add(X, Y)
	}
}

func BenchmarkInterfaceStrongExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var A Int = 4
		var B Int = 7

		AddI(A, B)
	}
}

func BenchmarkStructInterfaceStrongExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		A := S{4, 0}
		B := S{7, 0}

		AddS(A, B)
	}
}

func BenchmarkIntHeapPush(b *testing.B) {
	h := make(IntHeap, 0, SmallSize)
	H := &h

	for i := 0; i < b.N; i++ {
		if i%SmallSize == 0 {
			b.StopTimer()

			(*H) = (*H)[:0]
			heap.Init(H)

			b.StartTimer()
		}

		X := i / 2

		if i%2 == 0 {
			X = 3*i + 1
		}

		heap.Push(H, X)
	}
}

func BenchmarkStringHeapPush(b *testing.B) {
	h := make(StrHeap, 0, SmallSize)
	H := &h

	Hello := "hello"
	World := "world"

	for i := 0; i < b.N; i++ {
		if i%SmallSize == 0 {
			b.StopTimer()

			(*H) = (*H)[:0]
			heap.Init(H)

			b.StartTimer()
		}

		X := Hello

		if i%2 == 0 {
			X = World
		}

		heap.Push(H, X)
	}
}

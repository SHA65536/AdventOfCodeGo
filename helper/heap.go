package helper

import "cmp"

// For Max -> MakeHeap(MaxHeap)
// For Min -> MakeHeap(MinHeap)

// import "container/heap"

type Heap[T cmp.Ordered] struct {
	Values   []T
	LessFunc func(T, T) bool
}

func (h *Heap[T]) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap[T]) Len() int           { return len(h.Values) }
func (h *Heap[T]) Peek() T            { return h.Values[0] }
func (h *Heap[T]) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap[T]) Push(v interface{}) { h.Values = append(h.Values, v.(T)) }

func MakeHeap[T cmp.Ordered](less func(T, T) bool) *Heap[T] {
	return &Heap[T]{LessFunc: less}
}

func MaxHeap[T cmp.Ordered](i, j T) bool { return i > j }
func MinHeap[T cmp.Ordered](i, j T) bool { return i < j }

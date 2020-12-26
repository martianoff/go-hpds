package minheap

import (
	"github.com/maksimru/go-hpds/utils/arraylist"
	"github.com/maksimru/go-hpds/utils/comparator"
)

type Heap interface {
	IsValid() bool
	GetLength() int
	Top() interface{}
	Remove() interface{}
	Add(value interface{})
	IsEmpty() bool

	init()
	siftdown(index int)
	siftup(index int)
	leftIndexOf(index int) int
	rightIndexOf(index int) int
	parentIndexOf(index int) int
	swap(srcIndex int, dstIndex int)
}

type MinHeap struct {
	data       arraylist.ArrayList
	comparator comparator.Comparator
	length     int
}

// Check if heap is valid
// Return nil
// O(N) time, O(1) space
func (heap *MinHeap) IsValid() bool {
	for index := 1; index < heap.GetLength(); index++ {
		if heap.comparator.Greater(heap.data.Get(heap.parentIndexOf(index)), heap.data.Get(index)) {
			return false
		}
	}
	return true
}

// Get length of heap
// Return length
// O(1) time, O(1) space
func (heap MinHeap) GetLength() int {
	return heap.length
}

// Build Heap from ArrayList interface
// Return heap instance
// O(N) time, O(1) space
func NewMinHeap(arr arraylist.ArrayList, comp comparator.Comparator) *MinHeap {
	heap := new(MinHeap)
	if comp == nil {
		comp = comparator.NewIntComparator()
	}
	heap.data, heap.comparator, heap.length = arr, comp, arr.GetLength()
	heap.init()
	return heap
}

// Build Heap from array data source
// Return nil
// O(1) time, O(1) space
func (heap *MinHeap) init() {
	parent := (heap.GetLength() - 2) / 2
	for index := parent; index >= 0; index-- {
		heap.siftdown(index)
	}
}

// Moves value at index to proper heap position by sifting down
// Return nil
// O(log N) time, O(1) space
func (heap *MinHeap) siftdown(index int) {
	for heap.leftIndexOf(index) < heap.GetLength() {
		var idxSwap int
		if heap.rightIndexOf(index) < heap.GetLength() && heap.comparator.Less(heap.data.Get(heap.rightIndexOf(index)), heap.data.Get(heap.leftIndexOf(index))) {
			idxSwap = heap.rightIndexOf(index)
		} else {
			idxSwap = heap.leftIndexOf(index)
		}
		if heap.comparator.Greater(heap.data.Get(index), heap.data.Get(idxSwap)) {
			heap.swap(index, idxSwap)
			index = idxSwap
		} else {
			break
		}
	}
}

// Moves value at index to proper heap position by sifting up
// Return nil
// O(log N) time, O(1) space
func (heap *MinHeap) siftup(index int) {
	for index > 0 && heap.comparator.Less(heap.data.Get(index), heap.data.Get(heap.parentIndexOf(index))) {
		heap.swap(index, heap.parentIndexOf(index))
		index = heap.parentIndexOf(index)
	}
}

// Get index of left child
// Return index of left child
// O(1) time, O(1) space
func (heap *MinHeap) leftIndexOf(index int) int {
	return 2*index + 1
}

// Get index of right child
// Return index of right child
// O(1) time, O(1) space
func (heap *MinHeap) rightIndexOf(index int) int {
	return 2*index + 2
}

// Get index of parent
// Return index of parent
// O(1) time, O(1) space
func (heap *MinHeap) parentIndexOf(index int) int {
	return (index - 1) / 2
}

// Swap values in heap array
// Return nil
// O(1) time, O(1) space
func (heap *MinHeap) swap(srcIndex int, dstIndex int) {
	heap.data.Swap(srcIndex, dstIndex)
}

// Get top of the heap
// Return value on top
// O(1) time, O(1) space
func (heap *MinHeap) Top() interface{} {
	if heap.data.GetLength() == 0 {
		return nil
	}
	return heap.data.Get(0)
}

// Get top of the heap
// Return value on top
// O(log N) time, O(1) space
func (heap *MinHeap) Remove() interface{} {
	heap.swap(0, heap.GetLength()-1)
	heap.length--

	heap.siftdown(0)
	return heap.data.RemoveLast()
}

// Adds value to the heap
// Return nil
// O(log N) time, O(1) space
func (heap *MinHeap) Add(value interface{}) {
	heap.length++
	heap.data.Add(value)
	heap.siftup(heap.length - 1)
}

// Check if heap is empty
// Return true if heap is empty
// O(1) time, O(1) space
func (heap *MinHeap) IsEmpty() bool {
	return heap.length == 0
}

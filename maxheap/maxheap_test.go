package maxheap

import (
	"github.com/maksimru/go-hpds/utils/arraylist"
	"github.com/maksimru/go-hpds/utils/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getArray() [10]int {
	return [...]int{0, 5, 1, 6, 8, 3, 5, 9, 2, 6}
}

func TestMaxHeap_NewMaxHeap(t *testing.T) {
	arr := getArray()
	slice := arr[:]
	heap := NewMaxHeap(arraylist.NewIntArrayList(slice), nil)
	assert.Equal(t, 10, heap.GetLength())
	assert.Equal(t, 9, heap.Top())
	assert.Equal(t, true, heap.IsValid())
}

func TestMaxHeap_NewEmptyMaxHeap(t *testing.T) {
	heap := NewMaxHeap(arraylist.NewIntArrayList([]int{}), nil)
	assert.Equal(t, nil, heap.Top())
}

func TestMaxHeap_Remove(t *testing.T) {
	arr := getArray()
	slice := arr[:]
	heap := NewMaxHeap(arraylist.NewIntArrayList(slice), nil)
	assert.Equal(t, 9, heap.Remove())
	assert.Equal(t, 8, heap.Remove())
	assert.Equal(t, 6, heap.Remove())
	assert.Equal(t, 6, heap.Remove())
	assert.Equal(t, 5, heap.Remove())
	assert.Equal(t, 5, heap.Remove())
	assert.Equal(t, 3, heap.Remove())
	assert.Equal(t, 2, heap.Remove())
	assert.Equal(t, 1, heap.Remove())
	assert.Equal(t, 0, heap.Remove())
	assert.Equal(t, 0, heap.GetLength())
	assert.Equal(t, true, heap.IsEmpty())
	assert.Equal(t, true, heap.IsValid())
}

func TestMaxHeap_Add(t *testing.T) {
	arr := getArray()
	slice := arr[:]
	heap := NewMaxHeap(arraylist.NewIntArrayList(slice), nil)
	heap.Remove()
	assert.Equal(t, 9, heap.GetLength())
	assert.Equal(t, 8, heap.Top())
	heap.Add(10)
	assert.Equal(t, true, heap.IsValid())
	assert.Equal(t, 10, heap.Top())
	heap.Remove()
	assert.Equal(t, true, heap.IsValid())
	heap.Add(4)
	assert.Equal(t, 8, heap.Top())
	assert.Equal(t, true, heap.IsValid())
}

func TestMaxHeap_IsValid(t *testing.T) {
	arrSlice := []int{0, 5, 1, 6, 8, 3, 5, 9, 2, 6}
	heap := NewMaxHeap(arraylist.NewIntArrayList(arrSlice), comparator.NewIntComparator())
	assert.Equal(t, heap.IsValid(), true)
	arrSlice[3] = 232
	assert.Equal(t, heap.IsValid(), false)
}

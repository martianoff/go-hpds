package priorityqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getIntPrioritizedValues() IntPrioritizedValueList {
	arr := []IntPrioritizedValue{NewIntPrioritizedValue(21, 5), NewIntPrioritizedValue(2, 7), NewIntPrioritizedValue(432, 1)}
	return NewIntPrioritizedValueList(arr)
}

func TestPriorityQueue_NewPriorityQueueInt(t *testing.T) {
	pq := NewPriorityQueue(getIntPrioritizedValues(), NewIntPriorityComparator())
	assert.Equal(t, 3, pq.GetLength())
	assert.Equal(t, 2, pq.Top().GetValue())
	assert.Equal(t, 7, pq.Top().GetPriority())
}

func TestPriorityQueue_DequeueInt(t *testing.T) {
	pq := NewPriorityQueue(getIntPrioritizedValues(), NewIntPriorityComparator())
	item := pq.Dequeue()
	assert.Equal(t, 2, item.GetValue())
	assert.Equal(t, 7, item.GetPriority())
	assert.Equal(t, 2, pq.GetLength())
	assert.Equal(t, 21, pq.Top().GetValue())
	assert.Equal(t, 5, pq.Top().GetPriority())
	item = pq.Dequeue()
	assert.Equal(t, 21, item.GetValue())
	assert.Equal(t, 5, item.GetPriority())
	assert.Equal(t, 1, pq.GetLength())
	assert.Equal(t, 432, pq.Top().GetValue())
	assert.Equal(t, 1, pq.Top().GetPriority())
	item = pq.Dequeue()
	assert.Equal(t, 432, item.GetValue())
	assert.Equal(t, 1, item.GetPriority())
	assert.Equal(t, 0, pq.GetLength())
	assert.Equal(t, true, pq.IsEmpty())
}

func TestPriorityQueue_EnqueueInt(t *testing.T) {
	pq := NewPriorityQueue(getIntPrioritizedValues(), NewIntPriorityComparator())
	pq.Dequeue()
	pq.Dequeue()
	pq.Enqueue(NewIntPrioritizedValue(33, 44))
	assert.Equal(t, 2, pq.GetLength())
	assert.Equal(t, 33, pq.Top().GetValue())
	assert.Equal(t, 44, pq.Top().GetPriority())
	pq.Enqueue(NewIntPrioritizedValue(64, 1))
	assert.Equal(t, 3, pq.GetLength())
	assert.Equal(t, 33, pq.Top().GetValue())
	assert.Equal(t, 44, pq.Top().GetPriority())
}

func getFloat32PrioritizedValues() Float32PrioritizedValueList {
	arr := []Float32PrioritizedValue{NewFloat32PrioritizedValue(float32(21), 5), NewFloat32PrioritizedValue(float32(2), 7), NewFloat32PrioritizedValue(float32(432), 1)}
	return NewFloat32PrioritizedValueList(arr)
}

func TestPriorityQueue_NewPriorityQueueFloat32(t *testing.T) {
	pq := NewPriorityQueue(getFloat32PrioritizedValues(), NewFloat32PriorityComparator())
	assert.Equal(t, 3, pq.GetLength())
	assert.Equal(t, float32(2), pq.Top().GetValue())
	assert.Equal(t, 7, pq.Top().GetPriority())
}

func TestPriorityQueue_DequeueFloat32(t *testing.T) {
	pq := NewPriorityQueue(getFloat32PrioritizedValues(), NewFloat32PriorityComparator())
	item := pq.Dequeue()
	assert.Equal(t, float32(2), item.GetValue())
	assert.Equal(t, 7, item.GetPriority())
	assert.Equal(t, 2, pq.GetLength())
	assert.Equal(t, float32(21), pq.Top().GetValue())
	assert.Equal(t, 5, pq.Top().GetPriority())
	item = pq.Dequeue()
	assert.Equal(t, float32(21), item.GetValue())
	assert.Equal(t, 5, item.GetPriority())
	assert.Equal(t, 1, pq.GetLength())
	assert.Equal(t, float32(432), pq.Top().GetValue())
	assert.Equal(t, 1, pq.Top().GetPriority())
	item = pq.Dequeue()
	assert.Equal(t, float32(432), item.GetValue())
	assert.Equal(t, 1, item.GetPriority())
	assert.Equal(t, 0, pq.GetLength())
	assert.Equal(t, true, pq.IsEmpty())
}

func TestPriorityQueue_EnqueueFloat32(t *testing.T) {
	pq := NewPriorityQueue(getFloat32PrioritizedValues(), NewFloat32PriorityComparator())
	pq.Dequeue()
	pq.Dequeue()
	pq.Enqueue(NewFloat32PrioritizedValue(float32(33), 44))
	assert.Equal(t, 2, pq.GetLength())
	assert.Equal(t, float32(33), pq.Top().GetValue())
	assert.Equal(t, 44, pq.Top().GetPriority())
	pq.Enqueue(NewFloat32PrioritizedValue(float32(64), 1))
	assert.Equal(t, 3, pq.GetLength())
	assert.Equal(t, float32(33), pq.Top().GetValue())
	assert.Equal(t, 44, pq.Top().GetPriority())
}

func getFloat64PrioritizedValues() Float64PrioritizedValueList {
	arr := []Float64PrioritizedValue{NewFloat64PrioritizedValue(float64(21), 5), NewFloat64PrioritizedValue(float64(2), 7), NewFloat64PrioritizedValue(float64(432), 1)}
	return NewFloat64PrioritizedValueList(arr)
}

func TestPriorityQueue_NewPriorityQueueFloat64(t *testing.T) {
	pq := NewPriorityQueue(getFloat64PrioritizedValues(), NewFloat64PriorityComparator())
	assert.Equal(t, 3, pq.GetLength())
	assert.Equal(t, float64(2), pq.Top().GetValue())
	assert.Equal(t, 7, pq.Top().GetPriority())
}

func TestPriorityQueue_DequeueFloat64(t *testing.T) {
	pq := NewPriorityQueue(getFloat64PrioritizedValues(), NewFloat64PriorityComparator())
	item := pq.Dequeue()
	assert.Equal(t, float64(2), item.GetValue())
	assert.Equal(t, 7, item.GetPriority())
	assert.Equal(t, 2, pq.GetLength())
	assert.Equal(t, float64(21), pq.Top().GetValue())
	assert.Equal(t, 5, pq.Top().GetPriority())
	item = pq.Dequeue()
	assert.Equal(t, float64(21), item.GetValue())
	assert.Equal(t, 5, item.GetPriority())
	assert.Equal(t, 1, pq.GetLength())
	assert.Equal(t, float64(432), pq.Top().GetValue())
	assert.Equal(t, 1, pq.Top().GetPriority())
	item = pq.Dequeue()
	assert.Equal(t, float64(432), item.GetValue())
	assert.Equal(t, 1, item.GetPriority())
	assert.Equal(t, 0, pq.GetLength())
	assert.Equal(t, true, pq.IsEmpty())
}

func TestPriorityQueue_EnqueueFloat64(t *testing.T) {
	pq := NewPriorityQueue(getFloat64PrioritizedValues(), NewFloat64PriorityComparator())
	pq.Dequeue()
	pq.Dequeue()
	pq.Enqueue(NewFloat64PrioritizedValue(float64(33), 44))
	assert.Equal(t, 2, pq.GetLength())
	assert.Equal(t, float64(33), pq.Top().GetValue())
	assert.Equal(t, 44, pq.Top().GetPriority())
	pq.Enqueue(NewFloat64PrioritizedValue(float64(64), 1))
	assert.Equal(t, 3, pq.GetLength())
	assert.Equal(t, float64(33), pq.Top().GetValue())
	assert.Equal(t, 44, pq.Top().GetPriority())
}

func getStringPrioritizedValues() StringPrioritizedValueList {
	arr := []StringPrioritizedValue{NewStringPrioritizedValue("dog", 5), NewStringPrioritizedValue("cat", 7), NewStringPrioritizedValue("bird", 1)}
	return NewStringPrioritizedValueList(arr)
}

func TestPriorityQueue_NewPriorityQueueString(t *testing.T) {
	pq := NewPriorityQueue(getStringPrioritizedValues(), NewStringPriorityComparator())
	assert.Equal(t, 3, pq.GetLength())
	assert.Equal(t, "cat", pq.Top().GetValue())
	assert.Equal(t, 7, pq.Top().GetPriority())
}

func TestPriorityQueue_DequeueString(t *testing.T) {
	pq := NewPriorityQueue(getStringPrioritizedValues(), NewStringPriorityComparator())
	item := pq.Dequeue()
	assert.Equal(t, "cat", item.GetValue())
	assert.Equal(t, 7, item.GetPriority())
	assert.Equal(t, 2, pq.GetLength())
	assert.Equal(t, "dog", pq.Top().GetValue())
	assert.Equal(t, 5, pq.Top().GetPriority())
	item = pq.Dequeue()
	assert.Equal(t, "dog", item.GetValue())
	assert.Equal(t, 5, item.GetPriority())
	assert.Equal(t, 1, pq.GetLength())
	assert.Equal(t, "bird", pq.Top().GetValue())
	assert.Equal(t, 1, pq.Top().GetPriority())
	item = pq.Dequeue()
	assert.Equal(t, "bird", item.GetValue())
	assert.Equal(t, 1, item.GetPriority())
	assert.Equal(t, 0, pq.GetLength())
	assert.Equal(t, true, pq.IsEmpty())
}

func TestPriorityQueue_EnqueueString(t *testing.T) {
	pq := NewPriorityQueue(getStringPrioritizedValues(), NewStringPriorityComparator())
	pq.Dequeue()
	pq.Dequeue()
	pq.Enqueue(NewStringPrioritizedValue("snake", 44))
	assert.Equal(t, 2, pq.GetLength())
	assert.Equal(t, "snake", pq.Top().GetValue())
	assert.Equal(t, 44, pq.Top().GetPriority())
	pq.Enqueue(NewStringPrioritizedValue("elephant", 1))
	assert.Equal(t, 3, pq.GetLength())
	assert.Equal(t, "snake", pq.Top().GetValue())
	assert.Equal(t, 44, pq.Top().GetPriority())
}

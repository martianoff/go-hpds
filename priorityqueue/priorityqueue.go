package priorityqueue

import (
	"github.com/maksimru/go-hpds/maxheap"
	"github.com/maksimru/go-hpds/utils/arraylist"
	"github.com/maksimru/go-hpds/utils/comparator"
)

type PrioritizedValue interface {
	GetPriority() int
	GetValue() interface{}
}

type IntPrioritizedValue struct {
	value    int
	priority int
}

func NewIntPrioritizedValue(value int, priority int) IntPrioritizedValue {
	return IntPrioritizedValue{value: value, priority: priority}
}

func (v IntPrioritizedValue) GetValue() interface{} {
	return v.value
}

func (v IntPrioritizedValue) GetPriority() int {
	return v.priority
}

type IntPrioritizedValueList struct {
	data []IntPrioritizedValue
	arraylist.AbstractArrayList
}

func NewIntPrioritizedValueList(arr []IntPrioritizedValue) *IntPrioritizedValueList {
	list := IntPrioritizedValueList{arr, arraylist.AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = &list
	return &list
}

func (list *IntPrioritizedValueList) Get(index int) interface{} {
	return list.data[index]
}

func (list *IntPrioritizedValueList) Set(index int, value interface{}) {
	list.data[index] = value.(IntPrioritizedValue)
}

func (list *IntPrioritizedValueList) GetLength() int {
	return len(list.data)
}

func (list *IntPrioritizedValueList) RemoveLast() interface{} {
	valueForRemoval := list.data[list.GetLength()-1]
	list.data = list.data[0 : list.GetLength()-1]
	return valueForRemoval
}

func (list *IntPrioritizedValueList) Add(value interface{}) {
	list.data = append(list.data, value.(IntPrioritizedValue))
}

func (list *IntPrioritizedValueList) Clean() {
	list.data = nil
}

type IntPriorityComparator struct{ comparator.AbstractComparator }

func NewIntPriorityComparator() *IntPriorityComparator {
	cmp := IntPriorityComparator{comparator.AbstractComparator{}}
	cmp.AbstractComparator.Comparator = cmp
	return &cmp
}

func (cmp IntPriorityComparator) Less(value1 interface{}, value2 interface{}) bool {
	return value1.(IntPrioritizedValue).GetPriority() < value2.(IntPrioritizedValue).GetPriority()
}

func (cmp IntPriorityComparator) Equal(value1 interface{}, value2 interface{}) bool {
	return value1.(IntPrioritizedValue).GetPriority() == value2.(IntPrioritizedValue).GetPriority()
}

type Float32PrioritizedValue struct {
	value    float32
	priority int
}

func NewFloat32PrioritizedValue(value float32, priority int) Float32PrioritizedValue {
	return Float32PrioritizedValue{value: value, priority: priority}
}

func (v Float32PrioritizedValue) GetValue() interface{} {
	return v.value
}

func (v Float32PrioritizedValue) GetPriority() int {
	return v.priority
}

type Float32PrioritizedValueList struct {
	data []Float32PrioritizedValue
	arraylist.AbstractArrayList
}

func NewFloat32PrioritizedValueList(arr []Float32PrioritizedValue) *Float32PrioritizedValueList {
	list := Float32PrioritizedValueList{arr, arraylist.AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = &list
	return &list
}

func (list *Float32PrioritizedValueList) Get(index int) interface{} {
	return list.data[index]
}

func (list *Float32PrioritizedValueList) Set(index int, value interface{}) {
	list.data[index] = value.(Float32PrioritizedValue)
}

func (list *Float32PrioritizedValueList) GetLength() int {
	return len(list.data)
}

func (list *Float32PrioritizedValueList) RemoveLast() interface{} {
	valueForRemoval := list.data[list.GetLength()-1]
	list.data = list.data[0 : list.GetLength()-1]
	return valueForRemoval
}

func (list *Float32PrioritizedValueList) Add(value interface{}) {
	list.data = append(list.data, value.(Float32PrioritizedValue))
}

func (list *Float32PrioritizedValueList) Clean() {
	list.data = nil
}

type Float32PriorityComparator struct{ comparator.AbstractComparator }

func NewFloat32PriorityComparator() *Float32PriorityComparator {
	cmp := Float32PriorityComparator{comparator.AbstractComparator{}}
	cmp.AbstractComparator.Comparator = cmp
	return &cmp
}

func (cmp Float32PriorityComparator) Less(value1 interface{}, value2 interface{}) bool {
	return value1.(Float32PrioritizedValue).GetPriority() < value2.(Float32PrioritizedValue).GetPriority()
}

func (cmp Float32PriorityComparator) Equal(value1 interface{}, value2 interface{}) bool {
	return value1.(Float32PrioritizedValue).GetPriority() == value2.(Float32PrioritizedValue).GetPriority()
}

type Float64PrioritizedValue struct {
	value    float64
	priority int
}

func NewFloat64PrioritizedValue(value float64, priority int) Float64PrioritizedValue {
	return Float64PrioritizedValue{value: value, priority: priority}
}

func (v Float64PrioritizedValue) GetValue() interface{} {
	return v.value
}

func (v Float64PrioritizedValue) GetPriority() int {
	return v.priority
}

type Float64PrioritizedValueList struct {
	data []Float64PrioritizedValue
	arraylist.AbstractArrayList
}

func NewFloat64PrioritizedValueList(arr []Float64PrioritizedValue) *Float64PrioritizedValueList {
	list := Float64PrioritizedValueList{arr, arraylist.AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = &list
	return &list
}

func (list *Float64PrioritizedValueList) Get(index int) interface{} {
	return list.data[index]
}

func (list *Float64PrioritizedValueList) Set(index int, value interface{}) {
	list.data[index] = value.(Float64PrioritizedValue)
}

func (list *Float64PrioritizedValueList) GetLength() int {
	return len(list.data)
}

func (list *Float64PrioritizedValueList) RemoveLast() interface{} {
	valueForRemoval := list.data[list.GetLength()-1]
	list.data = list.data[0 : list.GetLength()-1]
	return valueForRemoval
}

func (list *Float64PrioritizedValueList) Add(value interface{}) {
	list.data = append(list.data, value.(Float64PrioritizedValue))
}

func (list *Float64PrioritizedValueList) Clean() {
	list.data = nil
}

type Float64PriorityComparator struct{ comparator.AbstractComparator }

func NewFloat64PriorityComparator() *Float64PriorityComparator {
	cmp := Float64PriorityComparator{comparator.AbstractComparator{}}
	cmp.AbstractComparator.Comparator = cmp
	return &cmp
}

func (cmp Float64PriorityComparator) Less(value1 interface{}, value2 interface{}) bool {
	return value1.(Float64PrioritizedValue).GetPriority() < value2.(Float64PrioritizedValue).GetPriority()
}

func (cmp Float64PriorityComparator) Equal(value1 interface{}, value2 interface{}) bool {
	return value1.(Float64PrioritizedValue).GetPriority() == value2.(Float64PrioritizedValue).GetPriority()
}

type StringPrioritizedValue struct {
	value    string
	priority int
}

func NewStringPrioritizedValue(value string, priority int) StringPrioritizedValue {
	return StringPrioritizedValue{value: value, priority: priority}
}

func (v StringPrioritizedValue) GetValue() interface{} {
	return v.value
}

func (v StringPrioritizedValue) GetPriority() int {
	return v.priority
}

type StringPrioritizedValueList struct {
	data []StringPrioritizedValue
	arraylist.AbstractArrayList
}

func NewStringPrioritizedValueList(arr []StringPrioritizedValue) *StringPrioritizedValueList {
	list := StringPrioritizedValueList{arr, arraylist.AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = &list
	return &list
}

func (list *StringPrioritizedValueList) Get(index int) interface{} {
	return list.data[index]
}

func (list *StringPrioritizedValueList) Set(index int, value interface{}) {
	list.data[index] = value.(StringPrioritizedValue)
}

func (list *StringPrioritizedValueList) GetLength() int {
	return len(list.data)
}

func (list *StringPrioritizedValueList) RemoveLast() interface{} {
	valueForRemoval := list.data[list.GetLength()-1]
	list.data = list.data[0 : list.GetLength()-1]
	return valueForRemoval
}

func (list *StringPrioritizedValueList) Add(value interface{}) {
	list.data = append(list.data, value.(StringPrioritizedValue))
}

func (list *StringPrioritizedValueList) Clean() {
	list.data = nil
}

type StringPriorityComparator struct{ comparator.AbstractComparator }

func NewStringPriorityComparator() *StringPriorityComparator {
	cmp := StringPriorityComparator{comparator.AbstractComparator{}}
	cmp.AbstractComparator.Comparator = cmp
	return &cmp
}

func (cmp StringPriorityComparator) Less(value1 interface{}, value2 interface{}) bool {
	return value1.(StringPrioritizedValue).GetPriority() < value2.(StringPrioritizedValue).GetPriority()
}

func (cmp StringPriorityComparator) Equal(value1 interface{}, value2 interface{}) bool {
	return value1.(StringPrioritizedValue).GetPriority() == value2.(StringPrioritizedValue).GetPriority()
}

type IPriorityQueue interface {
	Enqueue(PrioritizedValue)
	Dequeue() PrioritizedValue
	Top() PrioritizedValue
	GetLength() int
	IsEmpty() bool
}

type PriorityQueue struct {
	heap *maxheap.MaxHeap
}

// Build PriorityQueue from PrioritizedValueList
// Return priority queue instance
// O(N) time, O(1) space
func NewPriorityQueue(prioritylist arraylist.ArrayList, cmp comparator.Comparator) *PriorityQueue {
	pq := new(PriorityQueue)
	pq.heap = maxheap.NewMaxHeap(prioritylist, cmp)
	return pq
}

// Adds value to the priority queue
// Return nil
// O(log N) time, O(1) space
func (p PriorityQueue) Enqueue(value PrioritizedValue) {
	p.heap.Add(value)
}

// Extract value from top of the priority queue
// Return extracted value
// O(log N) time, O(1) space
func (p PriorityQueue) Dequeue() PrioritizedValue {
	return p.heap.Remove().(PrioritizedValue)
}

// Get top of the priority queue
// Return value on top
// O(1) time, O(1) space
func (p PriorityQueue) Top() PrioritizedValue {
	if p.heap.Top() == nil {
		return nil
	}
	return p.heap.Top().(PrioritizedValue)
}

// Get length of priority queue
// Return length
// O(1) time, O(1) space
func (p PriorityQueue) GetLength() int {
	return p.heap.GetLength()
}

// Check if priority queue is empty
// Return true if priority queue is empty
// O(1) time, O(1) space
func (p PriorityQueue) IsEmpty() bool {
	return p.heap.IsEmpty()
}

// Removes all elements from priority queue
// Return nil
// O(1) time, O(1) space
func (p PriorityQueue) Clean() {
	p.heap.Clean()
}

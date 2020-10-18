package queue

import "github.com/maksimru/go-hpds/doublylinkedlist"

type Queue struct {
	data *doublylinkedlist.DoublyLinkedList
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.data = doublylinkedlist.NewDoublyLinkedList()
	return queue
}

// Insert new value to queue
// Return nil
// O(1) time, O(1) space
func (queue *Queue) Enqueue(value interface{}) {
	queue.data.Append(value)
}

// Extract value from queue
// Return extracted value
// O(1) time, O(1) space
func (queue *Queue) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	return queue.data.Unshift().GetValue()
}

// Get top of queue
// Returns top value
// O(1) time, O(1) space
func (queue *Queue) Top() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	return queue.data.GetHead().GetValue()
}

// Get bottom of queue
// Returns bottom value
// O(1) time, O(1) space
func (queue *Queue) Bottom() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	return queue.data.GetTail().GetValue()
}

// Checks is queue empty
// Returns true or false
// O(1) time, O(1) space
func (queue *Queue) IsEmpty() bool {
	return queue.data.GetLength() == 0
}

// Get length of queue
// Returns length
// O(1) time, O(1) space
func (queue *Queue) GetLength() int {
	return queue.data.GetLength()
}

// Purge queue
// Returns nil
// O(1) time, O(1) space
func (queue *Queue) Purge() {
	queue.data.Purge()
}

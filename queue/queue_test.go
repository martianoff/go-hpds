package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	assert.Equal(t, 0, queue.GetLength())
}

func TestQueue_Enqueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	assert.Equal(t, 4, queue.GetLength())
	assert.Equal(t, 1, queue.Top())
	assert.Equal(t, 4, queue.Bottom())
}

func TestQueue_Dequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	assert.Equal(t, 1, queue.Dequeue())
	assert.Equal(t, 2, queue.Dequeue())
	assert.Equal(t, 3, queue.Dequeue())
	assert.Equal(t, 4, queue.Dequeue())
	assert.Equal(t, 0, queue.GetLength())
	assert.Nil(t, queue.Top())
	assert.Nil(t, queue.Bottom())
}

func TestQueue_Dequeue2(t *testing.T) {
	queue := NewQueue()
	assert.Nil(t, queue.Dequeue())
	assert.Nil(t, queue.Top())
	assert.Nil(t, queue.Bottom())
}

func TestQueue_Purge(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Purge()
	assert.Equal(t, 0, queue.GetLength())
	assert.Nil(t, queue.Top())
	assert.Nil(t, queue.Bottom())
}

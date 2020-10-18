package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	assert.Equal(t, 0, stack.GetLength())
}

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	assert.Equal(t, 4, stack.GetLength())
	assert.Equal(t, 4, stack.Top())
	assert.Equal(t, 1, stack.Bottom())
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	assert.Equal(t, 4, stack.Pop())
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, 0, stack.GetLength())
	assert.Nil(t, stack.Top())
	assert.Nil(t, stack.Bottom())
}

func TestStack_Pop2(t *testing.T) {
	stack := NewStack()
	assert.Nil(t, stack.Pop())
	assert.Nil(t, stack.Top())
	assert.Nil(t, stack.Bottom())
}

func TestStack_Purge(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Purge()
	assert.Equal(t, 0, stack.GetLength())
	assert.Nil(t, stack.Top())
	assert.Nil(t, stack.Bottom())
}

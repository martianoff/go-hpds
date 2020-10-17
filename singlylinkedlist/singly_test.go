package singlylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	var list = NewSinglyLinkedList()
	assert.Equal(t, 0, list.GetLength())
}

func TestSinglyLinkedList_Append(t *testing.T) {
	var list = NewSinglyLinkedList()
	var tail = list.Append(20)
	assert.Equal(t, 1, list.GetLength())
	assert.Equal(t, 20, tail.GetValue())
	assert.Equal(t, 20, list.GetHead().GetValue())
}

func TestSinglyLinkedList_Unshift(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	n := list.Unshift()
	assert.Equal(t, 1, list.GetLength())
	assert.Equal(t, 1, n.GetValue())
	assert.Equal(t, 2, list.GetHead().GetValue())
}

func TestSinglyLinkedList_Print(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	assert.Equal(t, "1 -> 2", list.Print())
}

func TestSinglyLinkedList_Append_Multiple(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	var tail = list.Append(4)
	assert.Equal(t, 4, list.GetLength())
	assert.Equal(t, 4, tail.GetValue())
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 4, list.GetTail().GetValue())
	assert.Equal(t, "1 -> 2 -> 3 -> 4", list.Print())
}

func TestSinglyLinkedList_Prepend(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	head := list.Prepend(5)
	assert.Equal(t, "5 -> 1 -> 2 -> 3", list.Print())
	assert.Equal(t, 5, list.GetHead().GetValue())
	assert.Equal(t, 5, head.GetValue())
	assert.Equal(t, 4, list.GetLength())
}

func TestSinglyLinkedList_Prepend2(t *testing.T) {
	var list = NewSinglyLinkedList()
	head := list.Prepend(1)
	assert.Equal(t, "1", list.Print())
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 1, head.GetValue())
	assert.Equal(t, 1, list.GetLength())
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	var head = list.Reverse()
	assert.Equal(t, 4, list.GetLength())
	assert.Equal(t, 4, head.GetValue())
	assert.Equal(t, 4, list.GetHead().GetValue())
	assert.Equal(t, 1, list.GetTail().GetValue())
	assert.Equal(t, "4 -> 3 -> 2 -> 1", list.Print())
}

func TestNode_AppendList(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	var list2 = NewSinglyLinkedList()
	list2.Append(5)
	list2.Append(6)
	list2.Append(7)
	list.GetHead().GetNext().AppendList(list2)
	assert.Equal(t, "1 -> 2 -> 5 -> 6 -> 7 -> 3", list.Print())
}

func TestNode_AppendList_End(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	var list2 = NewSinglyLinkedList()
	list2.Append(5)
	list2.Append(6)
	list2.Append(7)
	list.GetTail().AppendList(list2)
	assert.Equal(t, "1 -> 2 -> 3 -> 5 -> 6 -> 7", list.Print())
}

func TestNode_Append(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.GetTail().Append(3)
	list.GetTail().Append(4)
	list.GetHead().Append(2)
	assert.Equal(t, "1 -> 2 -> 3 -> 4", list.Print())
}

func TestNode_Append2(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(3)
	list.Append(4)
	list.GetTail().Append(2)
	assert.Equal(t, "1 -> 3 -> 4 -> 2", list.Print())
	assert.Equal(t, 2, list.GetTail().GetValue())
	assert.Equal(t, 4, list.GetLength())
}

func TestSinglyLinkedList_Join(t *testing.T) {
	var list = NewSinglyLinkedList()
	var list2 = NewSinglyLinkedList()
	list2.Append(5)
	list2.Append(6)
	list2.Append(7)
	list.Join(list2)
	assert.Equal(t, "5 -> 6 -> 7", list.Print())
}

func TestSinglyLinkedList_Purge(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Purge()
	assert.Nil(t, list.GetHead())
	assert.Nil(t, list.GetTail())
	assert.Equal(t, 0, list.GetLength())
	assert.Equal(t, "", list.Print())
}

func TestNode_Split(t *testing.T) {
	var list = NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	n := list.Append(5)
	list.Append(6)
	list.Append(7)
	list2 := n.Split()
	assert.Equal(t, "1 -> 2 -> 3 -> 5", list.Print())
	assert.Equal(t, "6 -> 7", list2.Print())
	assert.Equal(t, 4, list.GetLength())
	assert.Equal(t, 2, list2.GetLength())
	assert.Equal(t, 5, list.GetTail().GetValue())
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 7, list2.GetTail().GetValue())
	assert.Equal(t, 6, list2.GetHead().GetValue())
}

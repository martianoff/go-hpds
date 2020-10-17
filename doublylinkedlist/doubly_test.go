package doublylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	var list = NewDoublyLinkedList()
	assert.Equal(t, 0, list.GetLength())
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	var list = NewDoublyLinkedList()
	var tail = list.Insert(20)
	assert.Equal(t, 1, list.GetLength())
	assert.Equal(t, 20, tail.GetValue())
	assert.Equal(t, 20, list.GetHead().GetValue())
}

func TestDoublyLinkedList_Print(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(2)
	assert.Equal(t, "1 <-> 2", list.Print())
}

func TestDoublyLinkedList_Insert_Multiple(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	var tail = list.Insert(4)
	assert.Equal(t, 4, list.GetLength())
	assert.Equal(t, 4, tail.GetValue())
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 4, list.GetTail().GetValue())
	assert.Equal(t, "1 <-> 2 <-> 3 <-> 4", list.Print())
}

func TestDoublyLinkedList_Prepend(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	head := list.Prepend(5)
	assert.Equal(t, "5 <-> 1 <-> 2 <-> 3", list.Print())
	assert.Equal(t, 5, list.GetHead().GetValue())
	assert.Equal(t, 5, head.GetValue())
	assert.Equal(t, 4, list.GetLength())
}

func TestDoublyLinkedList_Prepend2(t *testing.T) {
	var list = NewDoublyLinkedList()
	head := list.Prepend(1)
	assert.Equal(t, "1", list.Print())
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 1, head.GetValue())
	assert.Equal(t, 1, list.GetLength())
}

func TestDoublyLinkedList_Reverse(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	var head = list.Reverse()
	assert.Equal(t, 4, list.GetLength())
	assert.Equal(t, 4, head.GetValue())
	assert.Equal(t, 4, list.GetHead().GetValue())
	assert.Equal(t, 1, list.GetTail().GetValue())
	assert.Equal(t, "4 <-> 3 <-> 2 <-> 1", list.Print())
}

func TestNode_AppendList(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	var list2 = NewDoublyLinkedList()
	list2.Insert(5)
	list2.Insert(6)
	list2.Insert(7)
	list.GetHead().GetNext().AppendList(list2)
	assert.Equal(t, "1 <-> 2 <-> 5 <-> 6 <-> 7 <-> 3", list.Print())
}

func TestNode_AppendList_End(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	var list2 = NewDoublyLinkedList()
	list2.Insert(5)
	list2.Insert(6)
	list2.Insert(7)
	list.GetTail().AppendList(list2)
	assert.Equal(t, "1 <-> 2 <-> 3 <-> 5 <-> 6 <-> 7", list.Print())
}

func TestNode_Iterations(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(3)
	list.Insert(4)
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 3, list.GetHead().GetNext().GetValue())
	assert.Equal(t, 4, list.GetHead().GetNext().GetNext().GetValue())
	assert.Equal(t, 3, list.GetHead().GetNext().GetNext().GetPrev().GetValue())
	assert.Equal(t, 1, list.GetHead().GetNext().GetNext().GetPrev().GetPrev().GetValue())
}

func TestNode_Append(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(3)
	list.Insert(4)
	list.GetHead().Append(2)
	assert.Equal(t, "1 <-> 2 <-> 3 <-> 4", list.Print())
}

func TestNode_Remove(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(3)
	n := list.Insert(4)
	list.Insert(5)
	n.Remove()
	assert.Equal(t, "1 <-> 3 <-> 5", list.Print())
	assert.Equal(t, 3, list.GetLength())
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 5, list.GetTail().GetValue())
}

func TestNode_Remove1(t *testing.T) {
	var list = NewDoublyLinkedList()
	n := list.Insert(1)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	n.Remove()
	assert.Equal(t, "3 <-> 4 <-> 5", list.Print())
	assert.Equal(t, 3, list.GetLength())
	assert.Equal(t, 3, list.GetHead().GetValue())
	assert.Equal(t, 5, list.GetTail().GetValue())
}

func TestNode_Remove2(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(3)
	list.Insert(4)
	n := list.Insert(5)
	n.Remove()
	assert.Equal(t, "1 <-> 3 <-> 4", list.Print())
	assert.Equal(t, 3, list.GetLength())
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 4, list.GetTail().GetValue())
}

func TestDoublyLinkedList_Join(t *testing.T) {
	var list = NewDoublyLinkedList()
	var list2 = NewDoublyLinkedList()
	list2.Insert(5)
	list2.Insert(6)
	list2.Insert(7)
	list.Join(list2)
	assert.Equal(t, "5 <-> 6 <-> 7", list.Print())
}

func TestDoublyLinkedList_Purge(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Purge()
	assert.Nil(t, list.GetHead())
	assert.Nil(t, list.GetTail())
	assert.Equal(t, 0, list.GetLength())
	assert.Equal(t, "", list.Print())
}

func TestNode_Split(t *testing.T) {
	var list = NewDoublyLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	n := list.Insert(5)
	list.Insert(6)
	list.Insert(7)
	list2 := n.Split()
	assert.Equal(t, "1 <-> 2 <-> 3 <-> 5", list.Print())
	assert.Equal(t, "6 <-> 7", list2.Print())
	assert.Equal(t, 4, list.GetLength())
	assert.Equal(t, 2, list2.GetLength())
	assert.Equal(t, 5, list.GetTail().GetValue())
	assert.Equal(t, 1, list.GetHead().GetValue())
	assert.Equal(t, 7, list2.GetTail().GetValue())
	assert.Equal(t, 6, list2.GetHead().GetValue())
}

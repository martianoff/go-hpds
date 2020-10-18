package stack

import "github.com/maksimru/go-hpds/doublylinkedlist"

type Stack struct {
	data *doublylinkedlist.DoublyLinkedList
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.data = doublylinkedlist.NewDoublyLinkedList()
	return stack
}

// Insert new value to stack
// Return nil
// O(1) time, O(1) space
func (stack *Stack) Push(value interface{}) {
	stack.data.Append(value)
}

// Pop value from stack
// Return extracted value
// O(1) time, O(1) space
func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data.Pop().GetValue()
}

// Get top of stack (last added)
// Returns top value
// O(1) time, O(1) space
func (stack *Stack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data.GetTail().GetValue()
}

// Get bottom of stack (first added)
// Returns bottom value
// O(1) time, O(1) space
func (stack *Stack) Bottom() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data.GetHead().GetValue()
}

// Checks is stack empty
// Returns true or false
// O(1) time, O(1) space
func (stack *Stack) IsEmpty() bool {
	return stack.data.GetLength() == 0
}

// Get length of stack
// Returns length
// O(1) time, O(1) space
func (stack *Stack) GetLength() int {
	return stack.data.GetLength()
}

// Purge stack
// Returns nil
// O(1) time, O(1) space
func (stack *Stack) Purge() {
	stack.data.Purge()
}

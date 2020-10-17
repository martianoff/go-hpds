package doublylinkedlist

import "fmt"

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return new(DoublyLinkedList)
}

// Insert new value to linkedlist
// Return inserted Node
// O(1) time, O(1) space
func (list *DoublyLinkedList) Append(value interface{}) *Node {
	node := NewNode(list, value, nil, list.tail)
	if list.tail == nil {
		list.tail, list.head = node, node
	} else {
		list.tail.next, list.tail = node, node
	}
	list.length++
	return list.tail
}

// Prepend new value to the begin of linkedlist
// Return inserted Node
// O(1) time, O(1) space
func (list *DoublyLinkedList) Prepend(value interface{}) *Node {
	node := NewNode(list, value, list.head, nil)
	if list.head == nil {
		list.tail, list.head = node, node
	} else {
		list.head.prev, list.head = node, node
	}
	list.length++
	return list.head
}

// Get head of linkedlist
// Returns head Node
// O(1) time, O(1) space
func (list *DoublyLinkedList) GetHead() *Node {
	return list.head
}

// Get tail of linkedlist
// Returns tail Node
// O(1) time, O(1) space
func (list *DoublyLinkedList) GetTail() *Node {
	return list.tail
}

// Get length of linkedlist
// Returns length
// O(1) time, O(1) space
func (list *DoublyLinkedList) GetLength() int {
	return list.length
}

// Purge linkedlist
// Returns nil
// O(1) time, O(1) space
func (list *DoublyLinkedList) Purge() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

// Reverse linked list
// Returns new head Node
// O(N) time, O(1) space
func (list *DoublyLinkedList) Reverse() *Node {
	var head = list.head
	list.head = list.head.ReverseFrom()
	list.tail = head
	return list.head
}

// Join other linked list to the end
// Returns new tail Node
// O(N) time, O(1) space
func (list *DoublyLinkedList) Join(other *DoublyLinkedList) *Node {
	if list.head == nil {
		list.head = other.head
		list.tail = other.tail
		list.length = other.length
	} else if other.head != nil {
		list.tail.next, other.head.prev = other.head, list.tail
		list.length += other.length
		list.tail = other.tail
	}
	return list.tail
}

// Print linked list
// Returns new head Node
// O(N) time, O(1) space
func (list *DoublyLinkedList) Print() string {
	var visualLinkedList string
	node := list.head
	for node != nil {
		visualLinkedList += fmt.Sprintf("%v", node.value)
		node = node.next
		if node != nil {
			visualLinkedList += " <-> "
		}
	}
	return visualLinkedList
}

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
	list  *DoublyLinkedList
}

//
func NewNode(list *DoublyLinkedList, value interface{}, next *Node, prev *Node) *Node {
	node := new(Node)
	node.list = list
	node.value = value
	node.next = next
	node.prev = prev
	return node
}

// Get next Node
// Returns value
// O(1) time, O(1) space
func (node *Node) GetNext() *Node {
	return node.next
}

// Get prev Node
// Returns value
// O(1) time, O(1) space
func (node *Node) GetPrev() *Node {
	return node.prev
}

// Get Node value
// Returns value
// O(1) time, O(1) space
func (node *Node) GetValue() interface{} {
	return node.value
}

// Get linkedlist from its node
// Returns linkedlist
// O(1) time, O(1) space
func (node *Node) GetList() *DoublyLinkedList {
	return node.list
}

// Append value after node
// Returns added node
// O(1) time, O(1) space
func (node *Node) Append(value interface{}) *Node {
	if node.next == nil {
		return node.GetList().Append(value)
	} else {
		newnode := NewNode(node.GetList(), value, node.next, node)
		node.next.prev, node.next = newnode, newnode
		node.GetList().length++
		return newnode
	}
}

// Prepend value before node
// Returns added node
// O(1) time, O(1) space
func (node *Node) Prepend(value interface{}) *Node {
	if node.prev == nil {
		return node.GetList().Prepend(value)
	} else {
		newnode := NewNode(node.GetList(), value, node, node.prev)
		node.prev.next, node.prev = newnode, newnode
		node.GetList().length++
		return newnode
	}
}

// Remove node
// Returns next node
// O(1) time, O(1) space
func (node *Node) Remove() *Node {
	switch {
	case node.next == nil:
		node.prev.next, node.GetList().tail = nil, node.prev
	case node.prev == nil:
		node.next.prev, node.GetList().head = nil, node.next
	default:
		node.prev.next, node.next.prev = node.next, node.prev
	}
	node.GetList().length--
	return node.next
}

// Append other linkedlist after node
// Returns next node
// O(1) time, O(1) space
func (node *Node) AppendList(list *DoublyLinkedList) *Node {
	if node.next != nil {
		list.tail.next, node.next = node.next, list.GetHead()
		node.GetList().length += list.length
	} else {
		node.GetList().Join(list)
	}
	return node.next
}

// Reverse part of linkedlist starting at Node N-K
// Returns new head Node
// O(K) time, O(1) space
func (node *Node) ReverseFrom() *Node {
	var tail, next = node, node
	var prev *Node
	for next != nil {
		tail = next.next
		next.next = prev
		prev, next = next, tail
		tail = prev
	}
	return tail
}

// Split linkedlist into two, next node become head of second list
// Returns new LinkedList
// O(K) time, O(1) space
func (node *Node) Split() *DoublyLinkedList {
	list := NewDoublyLinkedList()
	list.head = node.next
	list.head.prev = nil
	list.tail, node.GetList().tail = node.GetList().tail, node
	next := node.next
	for next != nil {
		list.length++
		next = next.next
	}
	node.next = nil
	node.GetList().length -= list.GetLength()
	return list
}

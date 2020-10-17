package singlylinkedlist

import "fmt"

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

// Insert new value to linkedlist
// Return inserted Node
// O(1) time, O(1) space
func (list *SinglyLinkedList) Append(value interface{}) *Node {
	node := NewNode(list, value, nil)
	if list.tail == nil {
		list.tail, list.head = node, node
	} else {
		list.tail.next = node
		list.tail = node
	}
	list.length++
	return list.tail
}

// Prepend new value to the begin of linkedlist
// Return inserted Node
// O(1) time, O(1) space
func (list *SinglyLinkedList) Prepend(value interface{}) *Node {
	node := NewNode(list, value, list.head)
	if list.head == nil {
		list.tail, list.head = node, node
	} else {
		list.head = node
	}
	list.length++
	return list.head
}

// Shift element from the head of linkedlist
// Return removed Node
// O(1) time, O(1) space
func (list *SinglyLinkedList) Unshift() *Node {
	head := list.GetHead()
	list.head = list.head.next
	list.length--
	return head
}

// Get head of linkedlist
// Returns head Node
// O(1) time, O(1) space
func (list *SinglyLinkedList) GetHead() *Node {
	return list.head
}

// Get tail of linkedlist
// Returns tail Node
// O(1) time, O(1) space
func (list *SinglyLinkedList) GetTail() *Node {
	return list.tail
}

// Get length of linkedlist
// Returns length
// O(1) time, O(1) space
func (list *SinglyLinkedList) GetLength() int {
	return list.length
}

// Purge linkedlist
// Returns nil
// O(1) time, O(1) space
func (list *SinglyLinkedList) Purge() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

// Reverse linked list
// Returns new head Node
// O(N) time, O(1) space
func (list *SinglyLinkedList) Reverse() *Node {
	var head = list.head
	list.head = list.head.ReverseFrom()
	list.tail = head
	return list.head
}

// Join other linked list to the end
// Returns new tail Node
// O(N) time, O(1) space
func (list *SinglyLinkedList) Join(other *SinglyLinkedList) *Node {
	if list.head == nil {
		list.head = other.head
		list.tail = other.tail
		list.length = other.length
	} else if other.head != nil {
		list.tail.next = other.head
		list.length += other.length
		list.tail = other.tail
	}
	return list.tail
}

// Print linked list
// Returns new head Node
// O(N) time, O(1) space
func (list *SinglyLinkedList) Print() string {
	var visualLinkedList string
	node := list.head
	for node != nil {
		visualLinkedList += fmt.Sprintf("%v", node.value)
		node = node.next
		if node != nil {
			visualLinkedList += " -> "
		}
	}
	return visualLinkedList
}

type Node struct {
	value interface{}
	next  *Node
	list  *SinglyLinkedList
}

//
func NewNode(list *SinglyLinkedList, value interface{}, next *Node) *Node {
	node := new(Node)
	node.list = list
	node.value = value
	node.next = next
	return node
}

// Get next Node
// Returns value
// O(1) time, O(1) space
func (node *Node) GetNext() *Node {
	return node.next
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
func (node *Node) GetList() *SinglyLinkedList {
	return node.list
}

// Append value after node
// Returns added node
// O(1) time, O(1) space
func (node *Node) Append(value interface{}) *Node {
	if node.next == nil {
		return node.GetList().Append(value)
	} else {
		node.next = NewNode(node.GetList(), value, node.next)
		node.GetList().length++
		return node.next
	}
}

// Append other linkedlist after node
// Returns next node
// O(1) time, O(1) space
func (node *Node) AppendList(list *SinglyLinkedList) *Node {
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
func (node *Node) Split() *SinglyLinkedList {
	list := NewSinglyLinkedList()
	list.head = node.next
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

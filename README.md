[![CircleCI](https://circleci.com/gh/maksimru/go-hpds.svg?style=svg&circle-token=1ff20ea621c63005a3dab9250eb68790c1b5049f)](https://circleci.com/gh/maksimru/go-hpds)
[![codecov](https://codecov.io/gh/maksimru/go-hpds/branch/master/graph/badge.svg?token=9R19KZFQ09)](https://codecov.io/gh/maksimru/go-hpds/)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/maksimru/go-hpds)](https://pkg.go.dev/github.com/maksimru/go-hpds)
[![Go Report Card](https://goreportcard.com/badge/github.com/maksimru/go-hpds)](https://goreportcard.com/report/github.com/maksimru/go-hpds)

# High-Performance Datastructures

## Supported Data Structures
- [SinglyLinkedList](#SinglyLinkedList)
- [DoublyLinkedList](#DoublyLinkedList)
- [Queue](#Queue)
- [Stack](#Stack)
- [MaxHeap](#MaxHeap)
- [MinHeap](#MinHeap)
- [PriorityQueue](#PriorityQueue)
- [Trie](#Trie)
- [Union-Find aka Disjoint-Set](#Union-find-Disjoint-set)

---
## SinglyLinkedList

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/singlylinkedlist"

var list = NewSinglyLinkedList()
list.Append(1)
list.Append(2)
list.Append(3)
var list2 = NewSinglyLinkedList()
list2.Append(5)
list2.Append(6)
list2.Append(7)
list.GetTail().AppendList(list2)
list.Print() //returns "1 -> 2 -> 3 -> 5 -> 6 -> 7"
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| SinglyLinkedList  | Purge  | O(1)  |  O(1)  | Purges entire list |
| SinglyLinkedList  | Join  | O(1)  |  O(1) | Joins two linkedlists |
| SinglyLinkedList  | GetLength  | O(1)  | O(1)  | Gets total length |
| SinglyLinkedList  | GetTail  | O(1)  | O(1)  | Gets tail node |
| SinglyLinkedList  | GetHead  | O(1)  | O(1)  | Gets head node |
| SinglyLinkedList  | Reverse | O(N)  |  O(1)  | Reverses entire list |
| SinglyLinkedList  | Append | O(1)  | O(1)  | Inserts element to the end |
| SinglyLinkedList  | Prepend | O(1)  | O(1)  | Inserts element to the begin |
| SinglyLinkedList  | Unshift | O(1)  | O(1)  | Shift element from the begin of linkedlist |
| SinglyLinkedList  | Print |  O(N)  | O(N)  | Prints linkedlist |
| Node (SinglyLinkedList)  | Split |  O(K)  | O(1)  | Splits linkedlist into two |
| Node (SinglyLinkedList)  | ReverseFrom |  O(K)  | O(1)  | Reverses part of linkedlist |
| Node (SinglyLinkedList)  | AppendList |  O(1)  | O(1)  | Appends other linkedlist after node |
| Node (SinglyLinkedList)  | Append |  O(1)  | O(1)  | Appends value after node |
| Node (SinglyLinkedList)  | GetList |  O(1)  | O(1)  | Returns list from node |
| Node (SinglyLinkedList)  | GetNext |  O(1)  | O(1)  | Gets next node |
| Node (SinglyLinkedList)  | GetValue |  O(1)  | O(1)  | Gets value of node |

---
## DoublyLinkedList

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/doublylinkedlist"

var list = NewDoublyLinkedList()
list.Append(1)
n := list.Append(2)
list.Append(3)
n.Remove()
list.Print() //returns "1 <-> 3"
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| DoublyLinkedList  | Purge  | O(1)  |  O(1)  | Purges entire list |
| DoublyLinkedList  | Join  | O(1)  |  O(1) | Joins two linkedlists |
| DoublyLinkedList  | GetLength  | O(1)  | O(1)  | Gets total length |
| DoublyLinkedList  | GetTail  | O(1)  | O(1)  | Gets tail node |
| DoublyLinkedList  | GetHead  | O(1)  | O(1)  | Gets head node |
| DoublyLinkedList  | Reverse | O(N)  |  O(1)  | Reverses entire list |
| DoublyLinkedList  | Append | O(1)  | O(1)  | Inserts element to the end |
| DoublyLinkedList  | Prepend | O(1)  | O(1)  | Inserts element to the begin |
| DoublyLinkedList  | Pop | O(1)  | O(1)  | Pop element from the end of linkedlist |
| DoublyLinkedList  | Unshift | O(1)  | O(1)  | Shift element from the begin of linkedlist |
| DoublyLinkedList  | Print |  O(N)  | O(N)  | Prints linkedlist |
| Node (DoublyLinkedList)  | Split |  O(K)  | O(1)  | Splits linkedlist into two |
| Node (DoublyLinkedList)  | ReverseFrom |  O(K)  | O(1)  | Reverses part of linkedlist |
| Node (DoublyLinkedList)  | AppendList |  O(1)  | O(1)  | Appends other linkedlist after node |
| Node (DoublyLinkedList)  | Append |  O(1)  | O(1)  | Appends value after node |
| Node (DoublyLinkedList)  | Prepend |  O(1)  | O(1)  | Prepends value before node |
| Node (DoublyLinkedList)  | Remove |  O(1)  | O(1)  | Removes node |
| Node (DoublyLinkedList)  | GetList |  O(1)  | O(1)  | Returns list from node |
| Node (DoublyLinkedList)  | GetNext |  O(1)  | O(1)  | Gets next node |
| Node (DoublyLinkedList)  | GetPrev |  O(1)  | O(1)  | Gets prev node |
| Node (DoublyLinkedList)  | GetValue |  O(1)  | O(1)  | Gets value of node |

---
## Queue

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/queue"

queue := NewQueue()
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)
queue.Enqueue(4)
for !queue.IsEmpty() {
    value := queue.Dequeue()
    // process the queue
}
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Queue  | Enqueue  | O(1)  |  O(1)  | Adds element to queue |
| Queue  | Dequeue  | O(1)  |  O(1) | Pops element from top of queue |
| Queue  | Top  | O(1)  | O(1)  | Returns top of queue (next candidate for dequeue) |
| Queue  | Bottom  | O(1)  | O(1)  | Returns bottom of queue (last candidate for dequeue) |
| Queue  | IsEmpty  | O(1)  | O(1)  | Checks queue is empty |
| Queue  | GetLength | O(1)  |  O(1)  | Gets queue length |
| Queue  | Purge | O(1)  |  O(1)  | Purges queue |

---
## Stack

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/stack"

stack := NewStack()
stack.Push(1)
stack.Push(2)
stack.Push(3)
stack.Push(4)
for !stack.IsEmpty() {
    value := stack.Pop()
    // process the stack
}
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Stack  | Push  | O(1)  |  O(1)  | Adds element to stack |
| Stack  | Pop  | O(1)  |  O(1) | Pops element from top of stack |
| Stack  | Top  | O(1)  | O(1)  | Get top of stack (last added) |
| Stack  | Bottom  | O(1)  | O(1)  | Returns bottom of stack (first added) |
| Stack  | IsEmpty  | O(1)  | O(1)  | Checks stack is empty |
| Stack  | GetLength | O(1)  |  O(1)  | Gets stack length |
| Stack  | Purge | O(1)  |  O(1)  | Purges stack |

---
## MaxHeap

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/maxheap"

arrSlice := []int{0,5,1,6,8,3,5,9,2,6}
heap := NewMaxHeap(arraylist.NewIntArrayList(arrSlice), comparator.NewIntComparator())
heap.Remove() //returns 9
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| MaxHeap  | Add  | O(log N)  |  O(1)  | Adds element to the heap |
| MaxHeap  | Remove  | O(log N)  |  O(1) | Removes element from the top of the heap |
| MaxHeap  | Top  | O(1)  | O(1)  | Get top of heap |
| MaxHeap  | IsValid  | O(N)  | O(1)  | Checks heap is valid1 |
| MaxHeap  | IsEmpty  | O(1)  | O(1)  | Checks heap is empty |
| MaxHeap  | GetLength | O(1)  |  O(1)  | Gets length of the heap |
| MaxHeap  | Clean | O(1)  |  O(1)  | Removes all elements from the heap |
| MaxHeap  | NewMaxHeap | O(N)  |  O(1)  | Build heap |
---

## MinHeap

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/minheap"

arrSlice := []int{0,5,1,6,8,3,5,9,2,6}
heap := NewMinHeap(arraylist.NewIntArrayList(arrSlice), comparator.NewIntComparator())
heap.Remove() //returns 0
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| MinHeap  | Add  | O(log N)  |  O(1)  | Adds element to the heap |
| MinHeap  | Remove  | O(log N)  |  O(1) | Removes element from the top of the heap |
| MinHeap  | Top  | O(1)  | O(1)  | Get top of heap |
| MinHeap  | IsValid  | O(N)  | O(1)  | Checks heap is valid1 |
| MinHeap  | IsEmpty  | O(1)  | O(1)  | Checks heap is empty |
| MinHeap  | GetLength | O(1)  |  O(1)  | Gets length of the heap |
| MinHeap  | Clean | O(1)  |  O(1)  | Removes all elements from the heap |
| MinHeap  | NewMinHeap | O(N)  |  O(1)  | Build heap |
---

## PriorityQueue

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/priorityqueue"

arr := []IntPrioritizedValue{NewIntPrioritizedValue(21,5),NewIntPrioritizedValue(2,7),NewIntPrioritizedValue(432,1)}
pq := NewPriorityQueue(NewIntPrioritizedValueList(arr), NewIntPriorityComparator())
for !pq.IsEmpty() {
    value := pq.Dequeue() //returns instance of IntPrioritizedValue value = 2, priority = 7
}
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| PriorityQueue  | Enqueue  | O(log N)  |  O(1)  | Adds element to the priority queue |
| PriorityQueue  | Dequeue  | O(log N)  |  O(1) | Removes element from the top of the priority queue |
| PriorityQueue  | Top  | O(1)  | O(1)  | Get top of priority queue |
| PriorityQueue  | IsEmpty  | O(1)  | O(1)  | Checks priority queue is empty |
| PriorityQueue  | GetLength | O(1)  |  O(1)  | Gets length of the priority queue |
| PriorityQueue  | Clean | O(1)  |  O(1)  | Removes all elements from the priority queue |
| PriorityQueue  | NewPriorityQueue | O(N)  |  O(1)  | Build priority queue |
---

## Trie

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/trie"

var trie = NewTrie()
trie.Add("word", 1)
trie.Add("work", 2)
trie.Add("wor", 3)
trie.Add("war", 4)
trie.Add("won", 5)
v, s := trie.SearchPattern("w?r") //s will return 3 and 4
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| -  | NewTrie  | O(1)  | O(1)  | Builds empty trie |
| Trie  | GetRoot  | O(1)  |  O(1)  | Returns root trie node |
| Trie  | Add  | O(M)  |  O(M) | Adds word with custom value to the trie |
| Trie  | Search  | O(M)  | O(1)  | Searches word in trie using exact match |
| Trie  | SearchPrefix  | O(M+N*L)  | O(1)  | Searches words in trie using prefix |
| Trie  | SearchPattern  | O(M*A^X)  | O(1)  | Searches keyword in Trie using pattern match, "?" will match any single char |
| -  | NewNode | O(1)  |  O(1)  | Creates node |
| Node (Trie)  | GetChar | O(1)  |  O(1)  | Get character in specific node |
| Node (Trie)  | GetValue | O(1)  |  O(1)  | Get value in specific node |
| Node (Trie)  | HasTerminator | O(1)  |  O(1)  | Checks if node is the end of word |
| Node (Trie)  | HasChildNodes | O(1)  |  O(1)  | Checks if node has sub nodes |
---

## Union-Find (Disjoint-Set)

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/unionfind"

union := NewUnionFind()
union.Union(1,2)
union.Union(2,3)
union.Union(4,3)
node := union.FindInSet(4)  //returns Node with value 1 and rank 4, meaning that all four nodes are in union
```

### Methods

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| -  | NewUnionFind  | O(1)  | O(1)  | Builds empty union-find |
| UnionFind  | FindInSet  | O(M)  |  O(1)  | Returns parent node of the union searching by value of its member |
| UnionFind  | Union  | O(M)  |  O(M) | Joins two unions using any values belonging to them |
| UnionFind  | Has | O(1)  |  O(1)  | Checks if specific union exists |
| -  | NewNode | O(1)  |  O(1)  | Creates node |
| Node (UnionFind)  | GetRank | O(1)  |  O(1)  | Get rank in specific node |
| Node (UnionFind)  | GetValue | O(1)  |  O(1)  | Get value of specific node |
---
## Testing

To run all tests in this module:

```
go test ./...
```

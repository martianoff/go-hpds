[![CircleCI](https://circleci.com/gh/maksimru/go-hpds.svg?style=svg&circle-token=1ff20ea621c63005a3dab9250eb68790c1b5049f)](https://circleci.com/gh/maksimru/go-hpds)
[![codecov](https://codecov.io/gh/maksimru/go-hpds/branch/master/graph/badge.svg?token=9R19KZFQ09)](https://codecov.io/gh/maksimru/go-hpds/)

# High-Performance Datastructures

## SinglyLinkedList:

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/singlylinkedlist"

var list = NewSinglyLinkedList()
list.Insert(1)
list.Insert(2)
list.Insert(3)
var list2 = NewSinglyLinkedList()
list2.Insert(5)
list2.Insert(6)
list2.Insert(7)
list.GetTail().AppendList(list2)
list.Print() //returns "1 -> 2 -> 3 -> 5 -> 6 -> 7"
```

### Functions

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| SinglyLinkedList  | Purge  | O(1)  |  O(1)  | Purges entire list |
| SinglyLinkedList  | Join  | O(1)  |  O(1) | Joins two linkedlists |
| SinglyLinkedList  | GetLength  | O(1)  | O(1)  | Gets total length |
| SinglyLinkedList  | GetTail  | O(1)  | O(1)  | Gets tail node |
| SinglyLinkedList  | GetHead  | O(1)  | O(1)  | Gets head node |
| SinglyLinkedList  | Reverse | O(N)  |  O(1)  | Reverses entire list |
| SinglyLinkedList  | Insert | O(1)  | O(1)  | Inserts element to the end |
| SinglyLinkedList  | Prepend | O(1)  | O(1)  | Inserts element to the begin |
| SinglyLinkedList  | Print |  O(N)  | O(N)  | Prints linkedlist |
| Node (SinglyLinkedList)  | Split |  O(K)  | O(1)  | Splits linkedlist into two |
| Node (SinglyLinkedList)  | ReverseFrom |  O(K)  | O(1)  | Reverses part of linkedlist |
| Node (SinglyLinkedList)  | AppendList |  O(1)  | O(1)  | Appends other linkedlist after node |
| Node (SinglyLinkedList)  | Append |  O(1)  | O(1)  | Appends value after node |
| Node (SinglyLinkedList)  | GetList |  O(1)  | O(1)  | Returns list from node |
| Node (SinglyLinkedList)  | GetNext |  O(1)  | O(1)  | Gets next node |
| Node (SinglyLinkedList)  | GetValue |  O(1)  | O(1)  | Gets value of node |

## DoublyLinkedList:

### Usage

```
import "github.com/maksimru/go-hpds"
import "github.com/maksimru/go-hpds/doublylinkedlist"

var list = DoublyLinkedList()
list.Insert(1)
n := list.Insert(2)
list.Insert(3)
n.Remove()
list.Print() //returns "1 <-> 3"
```

### Functions

|  Object  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| DoublyLinkedList  | Purge  | O(1)  |  O(1)  | Purges entire list |
| DoublyLinkedList  | Join  | O(1)  |  O(1) | Joins two linkedlists |
| DoublyLinkedList  | GetLength  | O(1)  | O(1)  | Gets total length |
| DoublyLinkedList  | GetTail  | O(1)  | O(1)  | Gets tail node |
| DoublyLinkedList  | GetHead  | O(1)  | O(1)  | Gets head node |
| DoublyLinkedList  | Reverse | O(N)  |  O(1)  | Reverses entire list |
| DoublyLinkedList  | Insert | O(1)  | O(1)  | Inserts element to the end |
| DoublyLinkedList  | Prepend | O(1)  | O(1)  | Inserts element to the begin |
| DoublyLinkedList  | Print |  O(N)  | O(N)  | Prints linkedlist |
| Node (DoublyLinkedList)  | Split |  O(K)  | O(1)  | Splits linkedlist into two |
| Node (DoublyLinkedList)  | ReverseFrom |  O(K)  | O(1)  | Reverses part of linkedlist |
| Node (DoublyLinkedList)  | AppendList |  O(1)  | O(1)  | Appends other linkedlist after node |
| Node (DoublyLinkedList)  | Append |  O(1)  | O(1)  | Appends value after node |
| Node (DoublyLinkedList)  | Remove |  O(1)  | O(1)  | Removes node |
| Node (DoublyLinkedList)  | GetList |  O(1)  | O(1)  | Returns list from node |
| Node (DoublyLinkedList)  | GetNext |  O(1)  | O(1)  | Gets next node |
| Node (DoublyLinkedList)  | GetPrev |  O(1)  | O(1)  | Gets prev node |
| Node (DoublyLinkedList)  | GetValue |  O(1)  | O(1)  | Gets value of node |

## Testing

To run all tests in this module:

```
go test ./...
```

[![CircleCI](https://circleci.com/gh/maksimru/go-linkedlist.svg?style=svg&circle-token=1ff20ea621c63005a3dab9250eb68790c1b5049f)](https://circleci.com/gh/maksimru/go-linkedlist)
[![codecov](https://codecov.io/gh/maksimru/go-linkedlist/branch/main/graph/badge.svg?token=9R19KZFQ09)](https://codecov.io/gh/maksimru/go-linkedlist/)

# High-Performance Linked List Datastructures

## Usage

```
import "github.com/maksimru/go-linkedlist"
import "github.com/maksimru/go-linkedlist/singlylinkedlist"

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

## SinglyLinkedList:

|  Datastructure  | Operation | Time Complexity  | Space Complexity   | Comment |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| SinglyLinkedList  | Purge  | O(1)  |  O(1)  | Purges entire list |
| SinglyLinkedList  | Join  | O(1)  |  O(1) | Joins two linkedlists |
| SinglyLinkedList  | GetLength  | O(1)  | O(1)  | Gets total length |
| SinglyLinkedList  | GetTail  | O(1)  | O(1)  | Gets tail node |
| SinglyLinkedList  | GetHead  | O(1)  | O(1)  | Gets head node |
| SinglyLinkedList  | Reverse | O(N)  |  O(1)  | Reverses entire list |
| SinglyLinkedList  | Insert | O(1)  | O(1)  | Insert element to the end |
| SinglyLinkedList  | Prepend | O(1)  | O(1)  | Insert element to the begin |
| SinglyLinkedList  | Print |  O(N)  | O(N)  | Prints linkedlist |

All packages in this module are importable by other modules, except for packages
located in the `internal` directory.

## Testing

To run all tests in this module:

```
go test ./...
```

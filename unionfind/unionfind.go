package unionfind

type UnionFind struct {
	data map[interface{}]*Node
}

func NewUnionFind() *UnionFind {
	union := UnionFind{data: make(map[interface{}]*Node)}
	return &union
}

// Get parent node of the union searching by value of its member
// Return union's parent Node
// O(m) time, O(1) space (in the worst case O(m), where m is the number of Union calls)
func (union *UnionFind) FindInSet(value interface{}) *Node {
	var node *Node
	var exists bool
	node, exists = union.data[value]
	if !exists {
		node = NewNode(value, 1, nil)
		node.parent = node
		union.data[value] = node
	}
	// path compression, remap child nodes directly to the head
	if node.parent != node {
		node.parent = union.FindInSet(node.parent.GetValue())
	}
	return node.parent
}

// Join two unions using any values belonging to them
// Return new union's parent Node
// O(m*a(n)) time, O(m) space (approximately O(m), where m is the number of calls, a(n) is the inverse Ackermann function)
func (union *UnionFind) Union(value1 interface{}, value2 interface{}) *Node {
	node1 := union.FindInSet(value1)
	node2 := union.FindInSet(value2)
	if node1 != node2 {
		if node1.GetRank() >= node2.GetRank() {
			node2.parent = node1
			node1.rank += node2.GetRank()
			return node1
		}
		node1.parent = node2
		node2.rank += node1.GetRank()
		return node2
	}
	return node1
}

type Node struct {
	parent *Node
	rank   int
	value  interface{}
}

func NewNode(value interface{}, rank int, parent *Node) *Node {
	node := new(Node)
	node.value = value
	node.rank = rank
	return node
}

// Get Node rank
// Returns rank
// O(1) time, O(1) space
func (node *Node) GetRank() int {
	return node.rank
}

// Get Node value
// Returns value
// O(1) time, O(1) space
func (node *Node) GetValue() interface{} {
	return node.value
}

package trie

import "github.com/maksimru/go-hpds/queue"

const wildcardChar = '?'

type ITrie interface {
	Search(string) (interface{}, bool)
	SearchPrefix(string) ([]interface{}, bool)
	SearchPattern(string) ([]interface{}, bool)
	Add(string, interface{})
	ExtractValues(*Node) []interface{}
}

type Trie struct {
	root *Node
}

// Build empty Trie
// Return Trie instance
// O(1) time, O(1) space
func NewTrie() Trie {
	return Trie{
		root: NewNode('_', nil),
	}
}

// Get Trie root
// Return root Trie node
// O(1) time, O(1) space
func (t Trie) GetRoot() *Node {
	return t.root
}

// Add word with some value to the Trie
// Return nil
// O(M) time, O(M) space (M is length of word)
func (t Trie) Add(word string, value interface{}) {
	var node = t.GetRoot()
	for _, char := range word {
		_, exists := node.ChildNodes[char]
		if !exists {
			node.ChildNodes[char] = NewNode(char, nil)
		}
		node = node.ChildNodes[char]
	}
	node.value = value
	node.terminator = true
}

// Search keyword in Trie using exact match
// Return stored value and search status true or false
// O(M) time, O(1) space (M is length of query)
func (t Trie) Search(query string) (interface{}, bool) {
	var node = t.GetRoot()
	for _, char := range query {
		_, exists := node.ChildNodes[char]
		if !exists {
			return nil, false
		}
		node = node.ChildNodes[char]
	}
	return node.value, node.terminator
}

// Search keyword in Trie using prefix match
// Return array of stored values and search status true or false
// O(M+N*L) time, O(N) space (M is length of prefix, N is number of results matched and L length of result word)
func (t Trie) SearchPrefix(prefix string) ([]interface{}, bool) {
	var node = t.GetRoot()
	for _, char := range prefix {
		_, exists := node.ChildNodes[char]
		if !exists {
			return nil, false
		}
		node = node.ChildNodes[char]
	}
	return t.ExtractValues(node), true
}

// Extract values starting from specific Node in Trie
// Return array of stored values
// Note: this function need to be customized for specific task in order to get optimal complexity
// O(N*L) time, O(N) space (N is number of results matched and L length of result word)
func (t Trie) ExtractValues(n *Node) []interface{} {
	var collector []interface{}
	q := queue.NewQueue()
	q.Enqueue(n)
	for !q.IsEmpty() {
		c := q.Dequeue().(*Node)
		if c.HasTerminator() {
			//TODO: optimize slice resizing
			collector = append(collector, c.GetValue())
		}
		for _, v := range c.ChildNodes {
			q.Enqueue(v)
		}
	}
	return collector
}

// Search keyword in Trie using pattern match, "?" will match any single char
// Return array of stored values and search status true or false
// Note: this function need to be customized for specific task in order to get optimal complexity
// O(M*(A^X)) time, O(N) space (M is length of pattern, A is number of chars in trie (in the most cases constant), X is the number of "?")
func (t Trie) SearchPattern(pattern string) ([]interface{}, bool) {
	q := queue.NewQueue()
	q.Enqueue(t.GetRoot())
	for _, char := range pattern {
		nq := queue.NewQueue()
		for !q.IsEmpty() {
			node := q.Dequeue().(*Node)
			if char == wildcardChar {
				for _, v := range node.ChildNodes {
					nq.Enqueue(v)
				}
			} else {
				_, exists := node.ChildNodes[char]
				if !exists && q.IsEmpty() {
					return nil, false
				} else if exists {
					nq.Enqueue(node.ChildNodes[char])
				}
			}
		}
		q = nq
	}
	var collector []interface{}
	for !q.IsEmpty() {
		c := q.Dequeue().(*Node)
		if c.HasTerminator() {
			//TODO: optimize slice resizing
			collector = append(collector, c.GetValue())
		}
	}
	return collector, true
}

type INode interface {
	GetValue() interface{}
	GetChar() rune
	HasTerminator() bool
	HasChildNodes() bool
}

type Node struct {
	value      interface{}
	ChildNodes map[rune]*Node
	char       rune
	terminator bool
}

func NewNode(char rune, value interface{}) *Node {
	return &Node{
		value:      value,
		ChildNodes: make(map[rune]*Node),
		char:       char,
		terminator: false,
	}
}

// Get Node char
// Returns character in specific node
// O(1) time, O(1) space
func (node Node) GetChar() rune {
	return node.char
}

// Get Node value
// Returns value
// O(1) time, O(1) space
func (node Node) GetValue() interface{} {
	return node.value
}

// Check if Node can be end of word
// Returns true if Node is the end of some word
// O(1) time, O(1) space
func (node Node) HasTerminator() bool {
	return node.terminator
}

// Check if Node has child nodes
// Returns true if Node has child nodes
// O(1) time, O(1) space
func (node Node) HasChildNodes() bool {
	return len(node.ChildNodes) > 0
}

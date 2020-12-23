package unionfind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	union := NewUnionFind()
	assert.IsType(t, UnionFind{}, *union)
}

func TestUnionFind_FindInSet(t *testing.T) {
	union := NewUnionFind()
	node := union.FindInSet(3)
	assert.IsType(t, Node{}, *node)
	assert.Equal(t, node, node.parent)
	assert.Equal(t, 1, node.GetRank())
	assert.Equal(t, 3, node.GetValue())
}

func TestUnionFind_Has(t *testing.T) {
	union := NewUnionFind()
	union.FindInSet(3)
	assert.Equal(t, true, union.Has(3))
	assert.Equal(t, false, union.Has(2))
}

func TestUnionFind_FindInSetTwoUnions(t *testing.T) {
	union := NewUnionFind()
	union.FindInSet(3)
	node := union.FindInSet(53)
	assert.IsType(t, Node{}, *node)
	assert.Equal(t, node, node.parent)
	assert.Equal(t, 1, node.GetRank())
	assert.Equal(t, 53, node.GetValue())
}

func TestUnionFind_FindInSetExisting(t *testing.T) {
	union := NewUnionFind()
	union.FindInSet(3)
	union.FindInSet(53)
	node := union.FindInSet(3)
	assert.IsType(t, Node{}, *node)
	assert.Equal(t, node, node.parent)
	assert.Equal(t, 1, node.GetRank())
	assert.Equal(t, 3, node.GetValue())
}

func TestUnionFind_FindInSetUnion(t *testing.T) {
	union := NewUnionFind()
	union.Union(1, 2)
	union.Union(2, 3)
	union.Union(4, 3)
	node := union.FindInSet(4)
	assert.Equal(t, 4, node.GetRank())
	assert.Equal(t, 1, node.GetValue())
}

func TestUnionFind_FindInSetMergeTwo(t *testing.T) {
	union := NewUnionFind()
	union.Union(1, 2)
	union.Union(2, 3)
	union.Union(4, 1)
	union.Union(11, 44)
	union.Union(22, 33)
	union.Union(44, 33)
	union.Union(22, 3)
	node := union.FindInSet(4)
	assert.Equal(t, 8, node.GetRank())
	assert.Equal(t, 11, node.GetValue())
}

func TestUnionFind_FindInSetMergeTwo2(t *testing.T) {
	union := NewUnionFind()
	union.Union(1, 2)
	union.Union(2, 3)
	union.Union(4, 1)
	union.Union(22, 33)
	union.Union(44, 33)
	union.Union(22, 3)
	node := union.FindInSet(4)
	assert.Equal(t, 7, node.GetRank())
	assert.Equal(t, 1, node.GetValue())
}

func TestUnionFind_FindInSetMergeTwo3(t *testing.T) {
	union := NewUnionFind()
	node := union.Union(1, 1)
	assert.Equal(t, 1, node.GetRank())
	assert.Equal(t, 1, node.GetValue())
}

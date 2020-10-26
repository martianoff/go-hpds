package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTrie(t *testing.T) {
	var trie = NewTrie()
	assert.IsType(t, Trie{}, trie)
}

func TestTrie_Add(t *testing.T) {
	var trie = NewTrie()
	trie.Add("hi", 1)
	trie.Add("hit", 2)
	trie.Add("ht", 6)
	h := trie.GetRoot().ChildNodes['h']
	assert.Equal(t, false, h.HasTerminator())
	assert.Equal(t, nil, h.GetValue())
	assert.Equal(t, 'h', h.GetChar())
	hi := h.ChildNodes['i']
	assert.Equal(t, true, hi.HasTerminator())
	assert.Equal(t, 1, hi.GetValue())
	assert.Equal(t, 'i', hi.GetChar())
	hit := hi.ChildNodes['t']
	assert.Equal(t, true, hit.HasTerminator())
	assert.Equal(t, 2, hit.GetValue())
	assert.Equal(t, 't', hit.GetChar())
	ht := h.ChildNodes['t']
	assert.Equal(t, true, ht.HasTerminator())
	assert.Equal(t, 6, ht.GetValue())
	assert.Equal(t, 't', ht.GetChar())
}

func TestTrie_Search(t *testing.T) {
	var trie = NewTrie()
	trie.Add("word", 1)
	trie.Add("hello", 2)
	trie.Add("wordy", 3)
	trie.Add("wor", 4)
	v, s := trie.Search("word")
	assert.Equal(t, 1, v)
	assert.Equal(t, true, s)
	v, s = trie.Search("hellow")
	assert.Equal(t, nil, v)
	assert.Equal(t, false, s)
	v, s = trie.Search("hello")
	assert.Equal(t, 2, v)
	assert.Equal(t, true, s)
	v, s = trie.Search("w")
	assert.Equal(t, nil, v)
	assert.Equal(t, false, s)
}

func TestTrie_SearchPrefix(t *testing.T) {
	var trie = NewTrie()
	trie.Add("word", 1)
	trie.Add("hello", 2)
	trie.Add("wordy", 3)
	trie.Add("wor", 4)
	v, s := trie.SearchPrefix("word")
	var r []interface{}
	r = append(r, 1, 3)
	assert.ElementsMatch(t, r, v)
	assert.Equal(t, true, s)
	v, s = trie.SearchPrefix("h")
	var r2 []interface{}
	r2 = append(r2, 2)
	assert.ElementsMatch(t, r2, v)
	assert.Equal(t, true, s)
	var r3 []interface{}
	v, s = trie.SearchPrefix("x")
	assert.ElementsMatch(t, r3, v)
	assert.Equal(t, false, s)
	v, s = trie.SearchPrefix("w")
	var r4 []interface{}
	r4 = append(r4, 4, 1, 3)
	assert.ElementsMatch(t, r4, v)
	assert.Equal(t, true, s)
}

func TestTrie_SearchPrefix2(t *testing.T) {
	var trie = NewTrie()
	trie.Add("word", "word")
	trie.Add("hello", "hello")
	trie.Add("wordy", "wordy")
	trie.Add("wor", "wor")
	v, s := trie.SearchPrefix("w")
	var r []interface{}
	r = append(r, "wor", "word", "wordy")
	assert.ElementsMatch(t, r, v)
	assert.Equal(t, true, s)
}

func TestNode_HasChildNodes(t *testing.T) {
	var trie = NewTrie()
	assert.Equal(t, false, trie.GetRoot().HasChildNodes())
	trie.Add("word", 1)
	assert.Equal(t, true, trie.GetRoot().HasChildNodes())
}

func TestTrie_SearchPattern(t *testing.T) {
	var trie = NewTrie()
	trie.Add("word", "word")
	trie.Add("work", "work")
	trie.Add("wor", "wor")
	trie.Add("war", "war")
	trie.Add("won", "won")
	v, s := trie.SearchPattern("w?r")
	var r []interface{}
	r = append(r, "wor", "war")
	assert.ElementsMatch(t, r, v)
	assert.Equal(t, true, s)
}

func TestTrie_SearchPattern2(t *testing.T) {
	var trie = NewTrie()
	trie.Add("wor", "wor")
	trie.Add("war", "war")
	trie.Add("won", "won")
	v, s := trie.SearchPattern("x??")
	var r []interface{}
	assert.ElementsMatch(t, r, v)
	assert.Equal(t, false, s)
}

func TestTrie_SearchPattern3(t *testing.T) {
	var trie = NewTrie()
	trie.Add("can", "can")
	trie.Add("pen", "pen")
	trie.Add("ten", "ten")
	v, s := trie.SearchPattern("??n")
	var r []interface{}
	r = append(r, "can", "pen", "ten")
	assert.ElementsMatch(t, r, v)
	assert.Equal(t, true, s)
}

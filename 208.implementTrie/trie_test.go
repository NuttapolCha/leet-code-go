package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")
	assert.Equal(t, true, trie.Search("apple"))
	assert.Equal(t, true, trie.StartsWith("apple"))
	assert.Equal(t, true, trie.StartsWith("app"))
	assert.Equal(t, false, trie.Search("app"))

	trie.Insert("app")
	assert.Equal(t, true, trie.Search("apple"))
	assert.Equal(t, true, trie.StartsWith("apple"))
	assert.Equal(t, true, trie.StartsWith("app"))
	assert.Equal(t, true, trie.Search("app"))
}

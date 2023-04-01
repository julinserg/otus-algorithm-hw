package p13prefixtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefixTree(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	require.Equal(t, true, trie.Search("apple"))
	require.Equal(t, false, trie.Search("app"))
	require.Equal(t, true, trie.StartsWith("app"))
	trie.Insert("app")
	require.Equal(t, true, trie.Search("app"))
}

// ["Trie","insert","search","search","search","startsWith","startsWith","startsWith"]
// [[],["hello"],["hell"],["helloa"],["hello"],["hell"],["helloa"],["hello"]]
// [null,null,false,false,true,true,false,true]
func TestPrefixTree2(t *testing.T) {
	trie := Constructor()
	trie.Insert("hello")
	require.Equal(t, false, trie.Search("hell"))
	require.Equal(t, false, trie.Search("helloa"))
	require.Equal(t, true, trie.Search("hello"))
	require.Equal(t, true, trie.StartsWith("hell"))
	require.Equal(t, false, trie.StartsWith("helloa"))
	require.Equal(t, true, trie.StartsWith("hello"))
}

// ["Trie","insert","insert","insert","insert","insert","insert","search","search","search","search","search","search","search"]
// [[],["app"],["apple"],["beer"],["add"],["jam"],["rental"],["apps"],["app"],["ad"],["applepie"],["rest"],["jan"]]
// [null,null,null,null,null,null,null,false,true,false,false,false,false]
func TestPrefixTree3(t *testing.T) {
	trie := Constructor()
	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("beer")
	trie.Insert("add")
	trie.Insert("jam")
	trie.Insert("rental")
	require.Equal(t, false, trie.Search("apps"))
	require.Equal(t, true, trie.Search("app"))
	require.Equal(t, false, trie.Search("ad"))
	require.Equal(t, false, trie.Search("applepie"))
	require.Equal(t, false, trie.Search("rest"))
	require.Equal(t, false, trie.Search("jan"))

}

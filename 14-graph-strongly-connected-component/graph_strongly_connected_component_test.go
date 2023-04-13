package p14graphstronglyconnectedcomponent

import (
	"testing"
)

func TestGraphPrint(t *testing.T) {
	var g Graph
	g.matrixLink = make(map[string]List, 0)
	g.matrixLink["A"] = NewList()
	g.matrixLink["A"].PushBack("B")
	g.matrixLink["A"].PushBack("E")
	g.matrixLink["B"] = NewList()
	g.matrixLink["B"].PushBack("A")
	g.matrixLink["B"].PushBack("C")
	g.matrixLink["C"] = NewList()
	g.matrixLink["C"].PushBack("D")
	g.matrixLink["D"] = NewList()
	g.matrixLink["D"].PushBack("C")
	g.matrixLink["D"].PushBack("F")
	g.matrixLink["E"] = NewList()
	g.matrixLink["E"].PushBack("F")

	g.Print("A")

	/*require.Equal(t, true, trie.Search("apple"))
	require.Equal(t, false, trie.Search("app"))
	require.Equal(t, true, trie.StartsWith("app"))
	trie.Insert("app")
	require.Equal(t, true, trie.Search("app"))*/
}

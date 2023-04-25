package p16graphminimumspanningtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphMinimumSpanningTree(t *testing.T) {
	graph := make([]NodeGraph, 0)
	graph = append(graph, NodeGraph{"A", []NodeParams{{"B", 1}, {"H", 2}, {"D", 6}}})
	graph = append(graph, NodeGraph{"B", []NodeParams{{"A", 1}, {"H", 3}, {"E", 4}}})
	graph = append(graph, NodeGraph{"C", []NodeParams{{"D", 1}, {"H", 5}, {"E", 4}, {"F", 1}, {"G", 2}}})
	graph = append(graph, NodeGraph{"D", []NodeParams{{"C", 1}, {"A", 6}, {"G", 4}}})
	graph = append(graph, NodeGraph{"E", []NodeParams{{"H", 1}, {"B", 4}, {"C", 4}, {"F", 3}}})
	graph = append(graph, NodeGraph{"F", []NodeParams{{"C", 1}, {"G", 3}, {"E", 3}}})
	graph = append(graph, NodeGraph{"G", []NodeParams{{"D", 4}, {"C", 2}, {"F", 3}}})
	graph = append(graph, NodeGraph{"H", []NodeParams{{"A", 2}, {"B", 3}, {"E", 1}, {"C", 5}}})

	gfscc := GraphMinimumSpanningTree{matrixLinkSrc: graph}

	resultTree := gfscc.MinimumSpanningTree()

	fmt.Println(resultTree)
	require.Equal(t, []string{"B", "H"}, resultTree["A"])
	require.Equal(t, []string{"A"}, resultTree["B"])
	require.Equal(t, []string{"D", "F", "G"}, resultTree["C"])
	require.Equal(t, []string{"C"}, resultTree["D"])
	require.Equal(t, []string{"H", "F"}, resultTree["E"])
	require.Equal(t, []string{"C", "E"}, resultTree["F"])
	require.Equal(t, []string{"C"}, resultTree["G"])
	require.Equal(t, []string{"E", "A"}, resultTree["H"])

}

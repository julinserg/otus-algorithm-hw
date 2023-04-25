package p17graphshortestpath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphMinimumSpanningTree(t *testing.T) {
	graph := make([]NodeGraph, 0)
	graph = append(graph, NodeGraph{"A", []NodeParams{{"B", 1}, {"C", 3}, {"E", 2}}})
	graph = append(graph, NodeGraph{"B", []NodeParams{{"A", 1}, {"D", 6}, {"C", 4}, {"G", 1}}})
	graph = append(graph, NodeGraph{"C", []NodeParams{{"A", 3}, {"B", 4}, {"D", 4}, {"E", 2}, {"F", 3}}})
	graph = append(graph, NodeGraph{"D", []NodeParams{{"B", 6}, {"C", 4}, {"F", 5}}})
	graph = append(graph, NodeGraph{"E", []NodeParams{{"A", 2}, {"C", 2}, {"G", 3}}})
	graph = append(graph, NodeGraph{"F", []NodeParams{{"G", 2}, {"C", 3}, {"D", 5}}})
	graph = append(graph, NodeGraph{"G", []NodeParams{{"B", 1}, {"E", 3}, {"F", 2}}})

	gsp := GraphShortestPath{matrixLinkSrc: graph}

	result := gsp.ShortestPath("B")

	fmt.Println(result)
	require.Equal(t, 1, result["A"])
	require.Equal(t, 0, result["B"])
	require.Equal(t, 4, result["C"])
	require.Equal(t, 6, result["D"])
	require.Equal(t, 3, result["E"])
	require.Equal(t, 3, result["F"])
	require.Equal(t, 1, result["G"])

}

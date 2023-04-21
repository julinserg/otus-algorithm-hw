package p15graphtopologicalsort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphSort(t *testing.T) {
	graph := make([]NodeGraph, 0)
	graph = append(graph, NodeGraph{"A", []string{"B"}})
	graph = append(graph, NodeGraph{"B", []string{"E"}})
	graph = append(graph, NodeGraph{"C", []string{"D"}})
	graph = append(graph, NodeGraph{"D", []string{"A", "B", "E", "F"}})
	graph = append(graph, NodeGraph{"E", []string{"G"}})
	graph = append(graph, NodeGraph{"F", []string{"H"}})
	graph = append(graph, NodeGraph{"G", []string{"H"}})
	graph = append(graph, NodeGraph{"H", []string{}})

	gfscc := GraphSort{matrixLinkSrc: graph}

	result := gfscc.Sort()
	fmt.Println(result)
	require.Equal(t, []string{"C", "D", "A", "F", "B", "E", "G", "H"}, result)

}

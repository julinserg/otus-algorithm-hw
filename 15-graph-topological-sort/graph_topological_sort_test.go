package p15graphtopologicalsort

import (
	"testing"
)

func listVertexName(list List) []string {
	result := make([]string, 0)
	if list == nil || list.Len() == 0 {
		return result
	}
	node := list.Front()
	for node != nil {
		result = append(result, node.Value.(string))
		node = node.Next
	}
	return result
}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

func TestGraphSort(t *testing.T) {
	graph := make([]NodeGraph, 0)
	graph = append(graph, NodeGraph{"A", []string{"B"}})
	graph = append(graph, NodeGraph{"B", []string{"A", "C"}})
	graph = append(graph, NodeGraph{"C", []string{"D"}})
	graph = append(graph, NodeGraph{"D", []string{"C"}})

	gfscc := GraphSort{matrixLinkSrc: graph}

	gfscc.Sort()
	gfscc.PrintMatrixAdj()
	/*fmt.Println("reachability from A")
	gfscc.PrintInvert("A")
	fmt.Println("reachability from D")
	gfscc.PrintInvert("D")
	comp1 := listVertexName(components[0])
	comp2 := listVertexName(components[1])
	fmt.Println(comp1)
	fmt.Println(comp2)
	require.Equal(t, true, sameStringSlice([]string{"A", "B"}, comp1))
	require.Equal(t, true, sameStringSlice([]string{"C", "D"}, comp2))*/

}

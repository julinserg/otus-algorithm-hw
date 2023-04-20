package p14graphstronglyconnectedcomponent

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestGraphComponents(t *testing.T) {
	graph := make(map[string]List, 0)
	graph["A"] = NewList()
	graph["A"].PushBack("B")
	graph["B"] = NewList()
	graph["B"].PushBack("A")
	graph["B"].PushBack("C")
	graph["C"] = NewList()
	graph["C"].PushBack("D")
	graph["D"] = NewList()
	graph["D"].PushBack("C")

	gfscc := GraphFindSCC{matrixLinkSrc: graph}

	components := gfscc.FindStronglyConnectedComponents()
	fmt.Println("reachability from A")
	gfscc.PrintInvert("A")
	fmt.Println("reachability from D")
	gfscc.PrintInvert("D")
	comp1 := listVertexName(components[0])
	comp2 := listVertexName(components[1])
	fmt.Println(comp1)
	fmt.Println(comp2)
	require.Equal(t, true, sameStringSlice([]string{"A", "B"}, comp1))
	require.Equal(t, true, sameStringSlice([]string{"C", "D"}, comp2))

}

func TestGraphComponentsComplex(t *testing.T) {
	graph := make(map[string]List, 0)
	graph["A"] = NewList()
	graph["B"] = NewList()
	graph["B"].PushBack("A")
	graph["C"] = NewList()
	graph["C"].PushBack("A")
	graph["C"].PushBack("E")
	graph["D"] = NewList()
	graph["D"].PushBack("C")
	graph["E"] = NewList()
	graph["E"].PushBack("B")
	graph["E"].PushBack("D")
	gfscc := GraphFindSCC{matrixLinkSrc: graph}

	components := gfscc.FindStronglyConnectedComponents()
	fmt.Println("reachability from A")
	gfscc.PrintInvert("A")
	fmt.Println("reachability from B")
	gfscc.PrintInvert("B")
	fmt.Println("reachability from C")
	gfscc.PrintInvert("C")
	fmt.Println("reachability from D")
	gfscc.PrintInvert("D")
	fmt.Println("reachability from E")
	gfscc.PrintInvert("E")
	comp1 := listVertexName(components[0])
	comp2 := listVertexName(components[1])
	comp3 := listVertexName(components[2])
	fmt.Println(comp1)
	fmt.Println(comp2)
	fmt.Println(comp3)
	require.Equal(t, true, sameStringSlice([]string{"C", "D", "E"}, comp1))
	require.Equal(t, true, sameStringSlice([]string{"B"}, comp2))
	require.Equal(t, true, sameStringSlice([]string{"A"}, comp3))

}
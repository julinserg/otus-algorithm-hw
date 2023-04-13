package p14graphstronglyconnectedcomponent

import "fmt"

type Graph struct {
	matrixLink map[string]List
}

func (g *Graph) Print(nameVertex string) {
	vertexVisit := make(map[string]string, 0)
	stack := NewList()
	stack.PushFront(nameVertex)
	for stack.Len() != 0 {
		vertex := stack.Front().Value.(string)
		stack.Remove(stack.Front())
		_, ok := vertexVisit[vertex]
		if ok {
			continue
		}
		vertexVisit[vertex] = vertex
		fmt.Println(vertex)
		listAdjVertex := g.matrixLink[vertex]
		if listAdjVertex == nil {
			continue
		}
		node := listAdjVertex.Front()
		for node != nil {
			stack.PushFront(node.Value)
			node = node.Next
		}
	}
}

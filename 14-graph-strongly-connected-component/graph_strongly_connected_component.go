package p14graphstronglyconnectedcomponent

import (
	"fmt"
	"sort"
)

type GraphFindSCC struct {
	matrixLink        map[string]List
	markRecursCounter int
}

type action func(string)

func (g *GraphFindSCC) dfs(nameVertex string, funcAction action) {
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
		funcAction(vertex)
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

func (g *GraphFindSCC) Print(nameBeginVertex string) {
	g.dfs(nameBeginVertex, func(vertex string) { fmt.Println(vertex) })
}

func isContains(list List, value string) bool {
	if list == nil || list.Len() == 0 {
		return false
	}
	node := list.Front()
	for node != nil {
		if node.Value.(string) == value {
			return true
		}
		node = node.Next
	}
	return false
}

func (g *GraphFindSCC) invert() {
	vertexVisit := make(map[string]string, 0)
	for key, _ := range g.matrixLink {
		_, ok := vertexVisit[key]
		if ok {
			continue
		}
		vertexVisit[key] = key
		node := g.matrixLink[key].Front()
		listVertexForRemove := make([]*ListItem, 0)
		for node != nil {
			vertex := node.Value.(string)
			if !isContains(g.matrixLink[vertex], key) {
				g.matrixLink[vertex].PushBack(key)
				listVertexForRemove = append(listVertexForRemove, node)
			}
			node = node.Next
		}
		for _, v := range listVertexForRemove {
			g.matrixLink[key].Remove(v)
		}
	}

}
func (g *GraphFindSCC) markLocal(markVertex map[string]int, vertexName string) {
	if len(markVertex) == len(g.matrixLink) {
		return
	}
	g.markRecursCounter++
	markVertex[vertexName] = g.markRecursCounter
	node := g.matrixLink[vertexName].Front()
	for node != nil {
		vertex := node.Value.(string)
		_, ok := markVertex[vertex]
		if ok {
			node = node.Next
			continue
		}
		g.markLocal(markVertex, vertex)
		g.markRecursCounter++
		markVertex[vertexName] = g.markRecursCounter
		node = node.Next
	}

}
func (g *GraphFindSCC) mark() []string {
	g.markRecursCounter = 0
	markVertex := make(map[string]int, 0)
	//---mark---
	for key := range g.matrixLink {
		g.markLocal(markVertex, key)
	}
	//---sort---
	sortVertex := make([]string, 0, len(markVertex))
	for key := range markVertex {
		sortVertex = append(sortVertex, key)
	}
	sort.SliceStable(sortVertex, func(i, j int) bool {
		return markVertex[sortVertex[i]] > markVertex[sortVertex[j]]
	})
	return sortVertex
}

func (g *GraphFindSCC) find(sortVertex []string) map[int]List {
	components := make(map[int]List, 0)
	vertexVisit := make(map[string]string, 0)
	numComponent := 0
	for _, vertex := range sortVertex {
		_, ok := vertexVisit[vertex]
		if ok {
			continue
		}
		vertexVisit[vertex] = vertex

		g.dfs(vertex, func(key string) {
			if components[numComponent] == nil {
				components[numComponent] = NewList()
			}
			components[numComponent].PushBack(key)
			vertexVisit[key] = key
		})
		numComponent++
	}
	return components
}

func (g *GraphFindSCC) FindStronglyConnectedComponents() map[int]List {
	sortVertex := g.mark()
	g.invert()
	components := g.find(sortVertex)
	return components
}

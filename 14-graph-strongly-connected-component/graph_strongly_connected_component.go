package p14graphstronglyconnectedcomponent

import "fmt"

type GraphFindSCC struct {
	matrixLink map[string]List
}

func (g *GraphFindSCC) printLocal(vertexVisit map[string]string, vertexName string) {
	vertexVisit[vertexName] = vertexName
	fmt.Println(vertexName)
	node := g.matrixLink[vertexName].Front()
	for node != nil {
		vertex := node.Value.(string)
		_, ok := vertexVisit[vertex]
		if ok {
			node = node.Next
			continue
		}
		g.printLocal(vertexVisit, vertex)
		node = node.Next
	}
}
func (g *GraphFindSCC) Print(nameBeginVertex string) {
	vertexVisit := make(map[string]string, 0)
	g.printLocal(vertexVisit, nameBeginVertex)
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
	for key := range g.matrixLink {
		_, ok := vertexVisit[key]
		if ok {
			continue
		}
		vertexVisit[key] = key
		node := g.matrixLink[key].Front()
		listVertexForRemove := make([]*ListItem, 0)
		for node != nil {
			vertex := node.Value.(string)
			_, ok := vertexVisit[vertex]
			if ok {
				node = node.Next
				continue
			}
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
func (g *GraphFindSCC) markLocal(sortVertex *[]string, vertexVisit map[string]string, vertexName string) {
	vertexVisit[vertexName] = vertexName
	node := g.matrixLink[vertexName].Front()
	for node != nil {
		vertex := node.Value.(string)
		_, ok := vertexVisit[vertex]
		if ok {
			node = node.Next
			continue
		}
		g.markLocal(sortVertex, vertexVisit, vertex)
		node = node.Next
	}
	*sortVertex = append(*sortVertex, vertexName)

}
func (g *GraphFindSCC) mark() []string {
	sortVertex := make([]string, 0)
	vertexVisit := make(map[string]string, 0)

	for key := range g.matrixLink {
		_, ok := vertexVisit[key]
		if ok {
			continue
		}
		g.markLocal(&sortVertex, vertexVisit, key)
	}
	return sortVertex
}

func (g *GraphFindSCC) findLocal(components map[int]List, numComponent int, vertexVisit map[string]string, vertexName string) {
	vertexVisit[vertexName] = vertexName
	if components[numComponent] == nil {
		components[numComponent] = NewList()
	}
	components[numComponent].PushBack(vertexName)
	node := g.matrixLink[vertexName].Front()
	for node != nil {
		vertex := node.Value.(string)
		_, ok := vertexVisit[vertex]
		if ok {
			node = node.Next
			continue
		}
		g.findLocal(components, numComponent, vertexVisit, vertex)
		node = node.Next
	}
}

func (g *GraphFindSCC) find(sortVertex []string) map[int]List {
	components := make(map[int]List, 0)
	vertexVisit := make(map[string]string, 0)
	numComponent := 0
	for i := len(sortVertex) - 1; i >= 0; i-- {
		vertex := sortVertex[i]
		_, ok := vertexVisit[vertex]
		if ok {
			continue
		}
		vertexVisit[vertex] = vertex
		g.findLocal(components, numComponent, vertexVisit, vertex)
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

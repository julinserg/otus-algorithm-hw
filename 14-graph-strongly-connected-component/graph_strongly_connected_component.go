package p14graphstronglyconnectedcomponent

import "fmt"

type GraphFindSCC struct {
	matrixLinkSrc    map[string]List
	matrixLinkInvert map[string]List
}

func (g *GraphFindSCC) printLocal(vertexVisit map[string]string, vertexName string) {
	vertexVisit[vertexName] = vertexName
	fmt.Println(vertexName)
	node := g.matrixLinkInvert[vertexName].Front()
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
func (g *GraphFindSCC) PrintInvert(nameBeginVertex string) {
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
	g.matrixLinkInvert = make(map[string]List, 0)
	for key := range g.matrixLinkSrc {
		node := g.matrixLinkSrc[key].Front()
		for node != nil {
			vertex := node.Value.(string)
			if g.matrixLinkInvert[vertex] == nil {
				g.matrixLinkInvert[vertex] = NewList()
			}
			g.matrixLinkInvert[vertex].PushBack(key)
			node = node.Next
		}
	}
}
func (g *GraphFindSCC) markLocal(sortVertex *[]string, vertexVisit map[string]string, vertexName string) {
	vertexVisit[vertexName] = vertexName
	node := g.matrixLinkSrc[vertexName].Front()
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

	for key := range g.matrixLinkSrc {
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
	node := g.matrixLinkInvert[vertexName].Front()
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

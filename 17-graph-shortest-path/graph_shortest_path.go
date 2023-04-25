package p17graphshortestpath

import (
	"fmt"
	"math"
)

type NodeParams struct {
	nodeName   string
	pathWeight int
}
type NodeGraph struct {
	nodeName        string
	listConnectNode []NodeParams
}

type GraphShortestPath struct {
	matrixLinkSrc      []NodeGraph
	mapNameNodeToIndex map[string]int
	matrixAdj          [][]int
}

func (g *GraphShortestPath) fillMatrixAdj() {
	g.mapNameNodeToIndex = make(map[string]int)
	g.matrixAdj = make([][]int, len(g.matrixLinkSrc))
	for i := range g.matrixAdj {
		g.matrixAdj[i] = make([]int, len(g.matrixLinkSrc))
	}

	for index, node := range g.matrixLinkSrc {
		g.mapNameNodeToIndex[node.nodeName] = index

	}
	for index, list := range g.matrixLinkSrc {
		for _, node := range list.listConnectNode {
			nodeIndex := g.mapNameNodeToIndex[node.nodeName]
			g.matrixAdj[index][nodeIndex] = node.pathWeight
		}
	}
}

func mapkey(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func (g *GraphShortestPath) PrintMatrixAdj() {
	fmt.Print("  ")
	for index := range g.matrixLinkSrc {
		key, _ := mapkey(g.mapNameNodeToIndex, index)
		fmt.Print(key, " ")
	}
	fmt.Print("\n")
	for row := 0; row < len(g.matrixAdj); row++ {
		key, _ := mapkey(g.mapNameNodeToIndex, row)
		fmt.Print(key, " ")
		for column := 0; column < len(g.matrixAdj); column++ {
			fmt.Print(g.matrixAdj[row][column], " ")
		}
		fmt.Print("\n")
	}
}

func (g *GraphShortestPath) findMinimal(array []int) int {
	minIndex := -1
	min := math.MaxInt32
	for i := 0; i < len(array); i++ {
		if array[i] == -1 {
			continue
		}
		if array[i] < min {
			min = array[i]
			minIndex = i
		}
	}
	return minIndex
}

func (g *GraphShortestPath) ShortestPath(startNodeName string) map[string]int {
	result := make(map[string]int, 0)
	g.fillMatrixAdj()
	g.PrintMatrixAdj()
	indexStart := g.mapNameNodeToIndex[startNodeName]
	markVertex := make([]int, len(g.matrixAdj))
	for i := 0; i < len(markVertex); i++ {
		if i != indexStart {
			markVertex[i] = math.MaxInt32
		}
	}

	for {
		minVertex := g.findMinimal(markVertex)
		if minVertex == -1 {
			break
		}
		for i := 0; i < len(g.matrixAdj); i++ {
			w := g.matrixAdj[minVertex][i]
			if w != 0 && markVertex[i] != -1 {
				if markVertex[minVertex]+w < markVertex[i] {
					markVertex[i] = markVertex[minVertex] + w
				}
			}
		}
		nodeName, _ := mapkey(g.mapNameNodeToIndex, minVertex)
		result[nodeName] = markVertex[minVertex]
		markVertex[minVertex] = -1
	}

	return result
}

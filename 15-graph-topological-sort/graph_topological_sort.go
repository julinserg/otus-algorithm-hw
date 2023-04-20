package p15graphtopologicalsort

import "fmt"

type NodeGraph struct {
	nodeName        string
	listConnectNode []string
}

type GraphSort struct {
	matrixLinkSrc      []NodeGraph
	mapNameNodeToIndex map[string]int
	matrixAdj          [][]int
}

func (g *GraphSort) fillMatrixAdj() {
	g.mapNameNodeToIndex = make(map[string]int)
	g.matrixAdj = make([][]int, len(g.matrixLinkSrc))
	for i := range g.matrixAdj {
		g.matrixAdj[i] = make([]int, len(g.matrixLinkSrc))
	}

	for index, node := range g.matrixLinkSrc {
		g.mapNameNodeToIndex[node.nodeName] = index

	}
	for index, list := range g.matrixLinkSrc {

		for _, nodeName := range list.listConnectNode {
			nodeIndex := g.mapNameNodeToIndex[nodeName]
			g.matrixAdj[index][nodeIndex] = 1
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

func (g *GraphSort) PrintMatrixAdj() {
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

func (g *GraphSort) Sort() {
	g.fillMatrixAdj()
}

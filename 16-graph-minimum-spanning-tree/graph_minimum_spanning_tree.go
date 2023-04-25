package p16graphminimumspanningtree

import (
	"fmt"
	"math"
)

type UnionFind struct {
	parents []int
}

func (u *UnionFind) create(size int) {
	u.parents = make([]int, size)
	for i := 0; i < len(u.parents); i++ {
		u.parents[i] = i
	}
}

func (u *UnionFind) find(index int) int {
	if u.parents[index] == index {
		return index
	}
	return u.find(u.parents[index])
}

func (u *UnionFind) union(index1 int, index2 int) {
	root1 := u.find(index1)
	root2 := u.find(index2)
	if root1 != root2 {
		u.parents[root1] = root2
	}
}

type NodeParams struct {
	nodeName   string
	pathWeight int
}
type NodeGraph struct {
	nodeName        string
	listConnectNode []NodeParams
}

type GraphMinimumSpanningTree struct {
	matrixLinkSrc      []NodeGraph
	mapNameNodeToIndex map[string]int
	matrixAdj          [][]int
	unionFind          UnionFind
}

func (g *GraphMinimumSpanningTree) fillMatrixAdj() {
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

func (g *GraphMinimumSpanningTree) PrintMatrixAdj() {
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

func (g *GraphMinimumSpanningTree) findMinimumEdge() (i int, j int) {
	indexI := -1
	indexJ := -1
	min := math.MaxInt32
	for i := 0; i < len(g.matrixAdj); i++ {
		for j := 0; j < len(g.matrixAdj); j++ {
			if g.matrixAdj[i][j] == 0 || g.matrixAdj[i][j] == math.MaxInt32 {
				continue
			}
			if g.matrixAdj[i][j] < min {
				min = g.matrixAdj[i][j]
				indexI = i
				indexJ = j
			}
		}
	}
	return indexI, indexJ
}

func (g *GraphMinimumSpanningTree) markDeleteEdge(i int, j int) {
	g.matrixAdj[i][j] = math.MaxInt32
	g.matrixAdj[j][i] = math.MaxInt32
}

func (g *GraphMinimumSpanningTree) MinimumSpanningTree() map[string][]string {
	g.fillMatrixAdj()
	g.PrintMatrixAdj()
	g.unionFind.create(len(g.matrixAdj))
	spanningTreeResult := make(map[string][]string, 0)
	for _, value := range g.matrixLinkSrc {
		spanningTreeResult[value.nodeName] = make([]string, 0)
	}
	for {
		i, j := g.findMinimumEdge()
		if i == -1 && j == -1 {
			break
		}
		g.markDeleteEdge(i, j)

		rootI := g.unionFind.find(i)
		rootJ := g.unionFind.find(j)
		if rootI == rootJ {
			continue
		}
		g.unionFind.union(i, j)

		node1, _ := mapkey(g.mapNameNodeToIndex, i)
		node2, _ := mapkey(g.mapNameNodeToIndex, j)
		spanningTreeResult[node1] = append(spanningTreeResult[node1], node2)
		spanningTreeResult[node2] = append(spanningTreeResult[node2], node1)

	}
	return spanningTreeResult
}

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

func (g *GraphSort) printSumRows(array []int, level int) {
	fmt.Printf("base sumrows level --- %d \n", level)
	fmt.Print("  ")
	for column := 0; column < len(g.matrixAdj); column++ {
		fmt.Print(array[column], " ")
	}
	fmt.Print("\n")
}

func (g *GraphSort) printSubNode(array []int, level int) {
	fmt.Printf("sub rows level --- %d \n", level)
	fmt.Print("  ")
	for column := 0; column < len(g.matrixAdj); column++ {
		fmt.Print(array[column], " ")
	}
	fmt.Print("\n")
}

func (g *GraphSort) sumRows() []int {
	sumRows := make([]int, len(g.matrixAdj))
	for i := 0; i < len(g.matrixAdj); i++ {
		for j := 0; j < len(g.matrixAdj); j++ {
			sumRows[j] += g.matrixAdj[i][j]
		}
	}
	return sumRows
}

func (g *GraphSort) findZeros(array []int) []int {
	result := make([]int, 0)
	for column := 0; column < len(array); column++ {
		if array[column] == 0 {
			result = append(result, column)
		}
	}
	return result
}

func (g *GraphSort) subtraction(array1 []int, array2 []int) []int {
	if len(array1) != len(array2) {
		panic("Error graph")
	}
	result := make([]int, len(array1))
	for i := 0; i < len(array1); i++ {
		if result[i] == -2 {
			continue
		}
		if array1[i] == 0 {
			result[i] = -1
		} else {
			result[i] = array1[i] - array2[i]
			if result[i] == 0 {
				result[i] = -2
			}
		}
	}
	return result
}

func (g *GraphSort) normalization(array []int) {
	for i := 0; i < len(array); i++ {
		if array[i] == -2 {
			array[i] = 0
		}
	}
}

func (g *GraphSort) Sort() []string {
	g.fillMatrixAdj()
	g.PrintMatrixAdj()
	sortNodeIndex := make([]int, 0)
	sumRow := g.sumRows()
	levelSumRow := sumRow
	for level := 0; level < len(g.matrixAdj)-1; level++ {
		g.printSumRows(levelSumRow, level)
		zeroIndexs := g.findZeros(levelSumRow)
		if len(zeroIndexs) == 0 {
			panic("Error graph")
		}
		sortNodeIndex = append(sortNodeIndex, zeroIndexs...)

		sub := levelSumRow
		for i := 0; i < len(zeroIndexs); i++ {
			g.printSubNode(g.matrixAdj[zeroIndexs[i]], level)
			sub = g.subtraction(sub, g.matrixAdj[zeroIndexs[i]])

		}
		g.normalization(sub)
		levelSumRow = sub
	}

	sortNodeName := make([]string, len(sortNodeIndex))
	for i := 0; i < len(sortNodeIndex); i++ {
		key, _ := mapkey(g.mapNameNodeToIndex, sortNodeIndex[i])
		sortNodeName[i] = key
	}
	return sortNodeName
}

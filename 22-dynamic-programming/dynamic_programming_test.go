package p22dp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirTree(t *testing.T) {
	tree := [][]int{
		{1, 0, 0, 0},
		{2, 3, 0, 0},
		{4, 5, 6, 0},
		{9, 8, 0, 3},
	}
	res := firTreeMaxSumPath(tree)
	require.Equal(t, 17, res)
}

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'0', '0', '0', '0', '0', '0'},
		{'1', '1', '0', '0', '0', '0'},
		{'1', '0', '0', '0', '0', '0'},
		{'1', '0', '0', '0', '0', '1'},
		{'0', '0', '0', '0', '0', '1'},
		{'0', '0', '1', '0', '0', '0'},
		{'1', '0', '0', '0', '1', '1'},
	}
	res := numIslands(grid)
	require.Equal(t, 5, res)

	grid2 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	res2 := numIslands(grid2)
	require.Equal(t, 1, res2)
}

func TestNod(t *testing.T) {
	require.Equal(t, 1, nod(5, 4))
	require.Equal(t, 6, nod(12, 6))
	require.Equal(t, 2, nod(2, 4))
	require.Equal(t, 1, nod(1, 5))
	require.Equal(t, 3, nod(3, 6))
	require.Equal(t, 4, nod(4, 12))
	require.Equal(t, 2, nod(6, 14))
	require.Equal(t, 14, nod(42, 56))
	require.Equal(t, 18, nod(461952, 116298))
	require.Equal(t, 32, nod(7966496, 314080416))
	require.Equal(t, 526, nod(24826148, 45296490))
}

func TestNodWithOptimize(t *testing.T) {
	require.Equal(t, 1, nodWithOptimize(5, 4))
	require.Equal(t, 6, nodWithOptimize(12, 6))
	require.Equal(t, 2, nodWithOptimize(2, 4))
	require.Equal(t, 1, nodWithOptimize(1, 5))
	require.Equal(t, 3, nodWithOptimize(3, 6))
	require.Equal(t, 4, nodWithOptimize(4, 12))
	require.Equal(t, 2, nodWithOptimize(6, 14))
	require.Equal(t, 14, nodWithOptimize(42, 56))
	require.Equal(t, 18, nodWithOptimize(461952, 116298))
	require.Equal(t, 32, nodWithOptimize(7966496, 314080416))
	require.Equal(t, 526, nodWithOptimize(24826148, 45296490))
}

func TestFiveNine(t *testing.T) {
	require.Equal(t, 2, fiveNine(1)) // 8 5
	require.Equal(t, 4, fiveNine(2)) // 58 85 55 88
	require.Equal(t, 6, fiveNine(3)) // 858 585 855 588 558 885
	require.Equal(t, 10, fiveNine(4))
}

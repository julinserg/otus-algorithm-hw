package p23hangar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHangarBruteForce(t *testing.T) {
	garden := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}
	square := hangarBruteForce(garden)
	require.Equal(t, 12, square)

	garden1 := [][]int{
		{0, 0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}
	square1 := hangarBruteForce(garden1)
	require.Equal(t, 8, square1)

	garden2 := [][]int{
		{0, 0, 0, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}
	square2 := hangarBruteForce(garden2)
	require.Equal(t, 8, square2)
}

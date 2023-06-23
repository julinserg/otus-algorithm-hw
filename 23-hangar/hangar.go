package p23hangar

import (
	"github.com/golang-collections/collections/stack"
)

func findLocalMaxSquare(garden [][]int, i, j int) int {
	maxSquare := 0
	h := 0
	wLimit := j + 1
	for i-h >= 0 && h <= i+1 && garden[i-h][j] != 1 {
		w := 0
		for j-w >= 0 && w <= j+1 && garden[i-h][j-w] != 1 {
			w++
		}
		if w < wLimit {
			wLimit = w
		}
		h++
		square := wLimit * h
		if square > maxSquare {
			maxSquare = square
		}
	}
	return maxSquare
}

func findLocalMaxSquareWithOptimization(garden [][]int, i, j int) int {
	maxSquare := 0
	h := 0
	wLimit := j + 1
	for i-h >= 0 && h <= i+1 && garden[i-h][j] != 1 {
		w := 0
		for w < wLimit && j-w >= 0 && w <= j+1 && garden[i-h][j-w] != 1 {
			w++
		}
		if w < wLimit {
			wLimit = w
		}
		if wLimit*(i+1) < maxSquare {
			break
		}
		h++
		square := wLimit * h
		if square > maxSquare {
			maxSquare = square
		}
	}
	return maxSquare
}

func calcColumnWidths(j int, garden [][]int, widths []int) []int {
	N := len(garden)
	for i := 0; i < N; i++ {
		if garden[i][j] == 1 {
			widths[i] = 0
		} else {
			widths[i]++
		}
	}
	return widths
}

func findLocalMaxSquareO3(garden [][]int, i, j int, widths []int) int {
	maxSquare := 0
	h := 0
	wLimit := j + 1
	for i-h >= 0 && h <= i+1 && garden[i-h][j] != 1 {
		w := widths[i-h]
		if w < wLimit {
			wLimit = w
		}
		if wLimit*(i+1) < maxSquare {
			break
		}
		h++
		square := wLimit * h
		if square > maxSquare {
			maxSquare = square
		}
	}
	return maxSquare
}

func hangarBruteForce(garden [][]int) int {
	N := len(garden)
	M := len(garden[0])
	maxSquare := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			square := findLocalMaxSquare(garden, i, j)
			if square > maxSquare {
				maxSquare = square
			}
		}
	}
	return maxSquare
}

func hangarBruteForceWithOptimization(garden [][]int) int {
	N := len(garden)
	M := len(garden[0])
	maxSquare := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			square := findLocalMaxSquareWithOptimization(garden, i, j)
			if square > maxSquare {
				maxSquare = square
			}
		}
	}
	return maxSquare
}

func hangarO3(garden [][]int) int {
	N := len(garden)
	M := len(garden[0])
	maxSquare := 0
	widths := make([]int, N)
	for j := 0; j < M; j++ {
		widths = calcColumnWidths(j, garden, widths)
		for i := 0; i < N; i++ {
			square := findLocalMaxSquareO3(garden, i, j, widths)
			if square > maxSquare {
				maxSquare = square
			}
		}
	}
	return maxSquare
}

func calcL(garden [][]int, L []int, widths []int) []int {
	stackL := stack.New()
	N := len(garden)
	for i := N - 1; i >= 0; i-- {
		for stackL.Len() != 0 {
			if widths[stackL.Peek().(int)] > widths[i] {
				L[stackL.Pop().(int)] = i + 1
			} else {
				break
			}
		}
		stackL.Push(i)
	}
	for stackL.Len() != 0 {
		L[stackL.Pop().(int)] = 0
	}
	return L
}

func calcR(garden [][]int, R []int, widths []int) []int {
	stackR := stack.New()
	N := len(garden)
	for i := 0; i < N; i++ {
		for stackR.Len() != 0 {
			if widths[stackR.Peek().(int)] > widths[i] {
				R[stackR.Pop().(int)] = i - 1
			} else {
				break
			}
		}
		stackR.Push(i)
	}
	for stackR.Len() != 0 {
		R[stackR.Pop().(int)] = N - 1
	}
	return R
}

func hangarO2(garden [][]int) int {
	N := len(garden)
	M := len(garden[0])
	maxSquare := 0
	widths := make([]int, N)
	L := make([]int, N)
	R := make([]int, N)
	for j := 0; j < M; j++ {
		widths = calcColumnWidths(j, garden, widths)
		L = calcL(garden, L, widths)
		R = calcR(garden, R, widths)
		for i := 0; i < N; i++ {
			w := widths[i]
			h := R[i] - L[i] + 1
			square := w * h
			if square > maxSquare {
				maxSquare = square
			}
		}
	}
	return maxSquare
}

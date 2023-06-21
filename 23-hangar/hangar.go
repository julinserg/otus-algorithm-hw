package p23hangar

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

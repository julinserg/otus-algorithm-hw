package p23hangar

/*func findMaxHeight(i, j int) int {
	maxHeight := 0
	return maxHeight
}

func findMaxWidth(i, j int) int {
	maxWidth := 0
	for i := 0; i < count; i++ {

	}
	return maxWidth
}*/

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

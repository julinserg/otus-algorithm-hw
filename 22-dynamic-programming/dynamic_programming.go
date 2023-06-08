package p22dp

func max(m1, m2 int) int {
	if m1 > m2 {
		return m1
	} else {
		return m2
	}
}

func firTreeMaxSumPath(tree [][]int) int {
	for i := len(tree) - 2; i >= 0; i-- {
		for j := 0; j < len(tree); j++ {
			if tree[i][j] == 0 {
				continue
			}
			tree[i][j] += max(tree[i+1][j], tree[i+1][j+1])
		}
	}
	return tree[0][0]
}

func removeIsland(grid [][]byte, i, j int) {
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	removeIsland(grid, i+1, j)
	removeIsland(grid, i-1, j)
	removeIsland(grid, i, j+1)
	removeIsland(grid, i, j-1)
}

func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				num++
				removeIsland(grid, i, j)
			}
		}
	}
	return num
}

func nod(number1, number2 int) int {
	if number1 == number2 {
		return number1
	} else if number1 > number2 {
		return nod(number1-number2, number2)
	} else {
		return nod(number2-number1, number1)
	}
}

func nodWithMultiple(number1, number2, multiple int) int {
	if number1 == number2 {
		return number1 * multiple
	} else if number1 > number2 {
		return nodWithMultiple(number1-number2, number2, multiple)
	} else {
		return nodWithMultiple(number2-number1, number1, multiple)
	}
}

func nodWithOptimize(number1, number2 int) int {
	if number1%2 == 0 && number2%2 == 0 {
		return nodWithMultiple(number1>>1, number2>>1, 2)
	} else if number1%2 != 0 && number2%2 == 0 {
		return nod(number1, number2>>1)
	} else if number1%2 == 0 && number2%2 != 0 {
		return nod(number1>>1, number2)
	} else {
		if number1 > number2 {
			return nod((number1-number2)>>1, number2)
		} else {
			return nod((number2-number1)>>1, number1)
		}
	}
}

func fiveNine(N int) int {
	x5 := 1
	x8 := 1
	x55 := 0
	x88 := 0

	for i := 2; i <= N; i++ {
		x5N := x8 + x88
		x55N := x5
		x8N := x5 + x55
		x88N := x8

		x5 = x5N
		x8 = x8N
		x55 = x55N
		x88 = x88N
	}
	return x5 + x8 + x55 + x88
}

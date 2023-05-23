package p181920substringsearch

func SearchSimple(str string, substr string) int {
	if len(str) == 0 && len(substr) == 0 {
		return 0
	}
	if len(str) < len(substr) {
		return -1
	}
	if len(substr) == 0 {
		return 0
	}

	for i := 0; i < len(str); i++ {
		indexStr := i
		indexSubStr := -1
		for j := 0; j < len(substr); j++ {
			if indexStr+j >= len(str) {
				break
			}
			if str[indexStr+j] != substr[j] {
				break
			}
			indexSubStr = j
		}
		if indexSubStr == len(substr)-1 {
			return i
		}

	}
	return -1
}

func createShiftTablePrefix(mask string) []int {
	shift := make([]int, 128)
	for i := 0; i < len(shift); i++ {
		shift[i] = len(mask)
	}
	for i := 0; i < len(mask)-1; i++ {
		shift[mask[i]] = len(mask) - i - 1
	}
	return shift
}

func SearchSimpleWithShiftPrefix(str string, substr string) int {
	if len(str) == 0 && len(substr) == 0 {
		return 0
	}
	if len(str) < len(substr) {
		return -1
	}
	if len(substr) == 0 {
		return 0
	}
	shift := createShiftTablePrefix(substr)
	i := 0
	for i <= len(str)-len(substr) {
		j := len(substr) - 1
		for j >= 0 {
			if i+j >= len(str) {
				break
			}
			if str[i+j] != substr[j] {
				break
			}
			j--
		}
		if j < 0 {
			return i
		}
		i += shift[str[i+len(substr)-1]]
	}
	return -1
}

// begin---- https://www-igm.univ-mlv.fr/~lecroq/string/node14.html#SECTION00140
func suffixes(substr string) []int {
	suff := make([]int, len(substr))
	m := len(substr)
	suff[m-1] = m
	g := m - 1
	f := 0
	for i := m - 2; i >= 0; i-- {
		if i > g && suff[i+m-1-f] < i-g {
			suff[i] = suff[i+m-1-f]
		} else {
			if i < g {
				g = i
			}
			f = i
			for g >= 0 && substr[g] == substr[g+m-1-f] {
				g--
			}
			suff[i] = f - g
		}
	}
	return suff
}

func createShiftTableSuffix(substr string) []int {
	table := make([]int, len(substr))
	suff := suffixes(substr)
	m := len(substr)
	for i := 0; i < m; i++ {
		table[i] = m
	}
	j := 0
	for i := m - 1; i >= 0; i-- {
		if suff[i] == i+1 {
			for ; j < m-1-i; j++ {
				if table[j] == m {
					table[j] = m - 1 - i
				}
			}
		}
	}
	for i := 0; i <= m-2; i++ {
		table[m-1-suff[i]] = m - 1 - i
	}
	return table
}

// end---- https://www-igm.univ-mlv.fr/~lecroq/string/node14.html#SECTION00140

func SearchSimpleWithShiftSuffix(str string, substr string) int {
	if len(str) == 0 && len(substr) == 0 {
		return 0
	}
	if len(str) < len(substr) {
		return -1
	}
	if len(substr) == 0 {
		return 0
	}
	shift := createShiftTableSuffix(substr)
	i := 0
	for i <= len(str)-len(substr) {
		j := len(substr) - 1
		for j >= 0 {
			if i+j >= len(str) {
				break
			}
			if str[i+j] != substr[j] {
				break
			}
			j--
		}
		if j < 0 {
			return i
		}
		i += shift[j]
	}
	return -1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func SearchBoyerMoore(str string, substr string) int {
	if len(str) == 0 && len(substr) == 0 {
		return 0
	}
	if len(str) < len(substr) {
		return -1
	}
	if len(substr) == 0 {
		return 0
	}
	shiftSufix := createShiftTableSuffix(substr)
	shiftPrefix := createShiftTablePrefix(substr)
	i := 0

	for i <= len(str)-len(substr) {
		j := len(substr) - 1
		for j >= 0 {
			if i+j >= len(str) {
				break
			}
			if str[i+j] != substr[j] {
				break
			}
			j--
		}
		if j < 0 {
			return i
		}
		i += max(shiftSufix[j], shiftPrefix[str[i+len(substr)-1]])
	}

	return -1
}

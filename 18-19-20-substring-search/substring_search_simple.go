package p181920substringsearch

func SearchSimple(str string, substr string) int {
	if len(str) == 0 && len(substr) == 0 {
		return 0
	}
	if len(str) < len(substr) {
		return -1
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

func createShiftTable(mask string) []int {
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
	shift := createShiftTable(substr)
	for i := 0; i < len(str); i += shift[str[i+len(substr)-1]] {
		indexStr := i
		indexSubStr := len(substr)
		for j := len(substr) - 1; j >= 0; j-- {
			if indexStr+j >= len(str) {
				break
			}
			if str[indexStr+j] != substr[j] {
				break
			}
			indexSubStr = j
		}
		if indexSubStr <= 0 {
			return i
		}
	}
	return -1
}

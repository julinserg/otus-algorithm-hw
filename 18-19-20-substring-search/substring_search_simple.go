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

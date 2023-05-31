package p181920substringsearch

var alphabet = map[int]string{
	0: "A",
	1: "B",
	2: "C",
}

func CreateStateMachine(pattern string) [][]int {
	delta := make([][]int, len(pattern))
	for i := range delta {
		delta[i] = make([]int, len(alphabet))
	}
	for q := 0; q < len(pattern); q++ {
		for a := 0; a < len(alphabet); a++ {
			line := pattern[:q] + alphabet[a]
			k := q + 1
			for pattern[:k] != line[len(line)-k:] {
				k--
			}
			delta[q][a] = k
		}
	}
	return delta
}

func mapkey(m map[int]string, value string) (key int, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func SearchByStateMachine(text string, delta [][]int) int {
	q := 0
	for i := 0; i < len(text); i++ {
		key, ok := mapkey(alphabet, string(text[i]))
		if !ok {
			continue
		}
		q = delta[q][key]
		if q == len(delta) {
			return i - len(delta) + 1
		}
	}
	return -1
}

func CreatePiSlow(pattern string) []int {
	pi := make([]int, len(pattern)+1)
	for q := 0; q <= len(pattern); q++ {
		line := pattern[:q]
		for length := 0; length < q; length++ {
			if line[:length] == line[len(line)-length:] {
				pi[q] = length
			}
		}
	}
	return pi
}

func CreatePiFast(pattern string) []int {
	pi := make([]int, len(pattern)+1)
	pi[1] = 0
	for q := 1; q < len(pattern); q++ {
		length := pi[q]
		for length > 0 && pattern[length] != pattern[q] {
			length = pi[length]
		}
		if pattern[length] == pattern[q] {
			length++
		}
		pi[q+1] = length
	}
	return pi
}

func SearchKnuthMorrisPratt(str string, substr string) int {
	if len(str) == 0 && len(substr) == 0 {
		return 0
	}
	if len(str) < len(substr) {
		return -1
	}
	if len(substr) == 0 {
		return 0
	}
	pi := CreatePiFast(substr)
	j := 0
	for i := 0; i < len(str); i++ {
		for (j > 0) && (str[i] != substr[j]) {
			j = pi[j-1]
		}
		if str[i] == substr[j] {
			j++
		}
		if j == len(substr) {
			return i - len(substr) + 1
		}
	}
	return -1
}

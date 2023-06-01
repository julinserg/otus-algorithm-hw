package p21rle

import (
	"strconv"
)

func RLEEncode(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		count := 1
		for i+1 < len(str) && str[i] == str[i+1] {
			count++
			i++
		}
		result += string(str[i]) + strconv.Itoa(count)
	}
	return result
}

func RLEDecode(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		num, err := strconv.Atoi(string(str[i]))
		if err == nil {
			for j := 0; j < num; j++ {
				result += string(str[i-1])
			}
		}
	}
	return result
}

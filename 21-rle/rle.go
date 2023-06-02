package p21rle

func RleEncode(data []byte) []byte {
	result := make([]byte, 0)
	for i := 0; i < len(data); i++ {
		count := 1
		for i+1 < len(data) && data[i] == data[i+1] && count < 64 {
			count++
			i++
		}
		if count > 1 || data[i] >= 192 {
			result = append(result, 191+byte(count))
			result = append(result, data[i])
		} else {
			result = append(result, data[i])
		}
	}
	return result
}

func RleDecode(data []byte) []byte {
	result := make([]byte, 0)
	for i := 0; i < len(data); i++ {
		if data[i] < 192 {
			result = append(result, data[i])
		} else if i+1 < len(data) {
			for j := 0; j < int(data[i])-191; j++ {
				result = append(result, data[i+1])
			}
			i++
		}
	}
	return result
}

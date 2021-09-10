package util

// IsContainInt64 ...
func IsContainInt64(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// IsContainString ...
func IsContainString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// RemoveDupInt64 ...
func RemoveDupInt64(items []int64) []int64 {
	result := make([]int64, 0, len(items))
	tempMap := map[int64]byte{}
	for _, eachItem := range items {
		l := len(tempMap)
		tempMap[eachItem] = 0
		if len(tempMap) != l {
			result = append(result, eachItem)
		}
	}
	return result
}

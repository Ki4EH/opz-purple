package utils

import (
	"sort"
)

func GetKeys[T any](m map[string]T) []string {
	var result []string

	for key := range m {
		result = append(result, key)
	}

	sort.Strings(result)

	return result
}

func GetParentID(value int64, data []int64) int64 {
	if value == 1 {
		return 1
	}

	left, right := 0, len(data)-1

	for left != right-1 {
		mid := left + (right-left)/2
		if data[mid] > value {
			right = mid
		} else if data[mid] < value {
			left = mid
		} else {
			return 1
		}
	}
	switch {
	case value == data[left] || value == data[right]:
		return 1
	case value < data[right]:
		return data[left]
	default:
		return data[right]
	}
}

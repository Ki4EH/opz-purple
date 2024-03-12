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

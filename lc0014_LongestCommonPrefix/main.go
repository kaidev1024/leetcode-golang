package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	nStrings := len(strs)
	if nStrings == 0 {
		return ""
	}
	if nStrings == 1 {
		return strs[0]
	}

	capacity := len(strs[0])

	var result strings.Builder
	result.Grow(capacity)

	for j := 0; j < capacity; j++ {

		for i := 1; i < nStrings; i++ {
			if len(strs[i]) == j {
				return result.String()
			}
			if strs[i][j] != strs[0][j] {
				return result.String()
			}
		}
		result.WriteByte(strs[0][j])
	}
	return result.String()
}

func main() {
	strs := []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(strs)

	fmt.Println("the longest common prefix is ", result)
}

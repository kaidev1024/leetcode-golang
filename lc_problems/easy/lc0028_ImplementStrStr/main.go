package main

import "fmt"

func strStr(haystack string, needle string) int {
	nh := len(haystack)
	nn := len(needle)

	if nn == 0 {
		return 0
	}

	if nh < nn {
		return -1
	}

	if nh == nn {
		if haystack == needle {
			return 0
		}
		return -1
	}

	needleIndices := make([]int, nn)
	l := 0

	for i := 1; i < nn; {
		if needle[i] == needle[l] {
			l++
			needleIndices[i] = l
			i++
		} else {
			if l == 0 {
				i++
			} else {
				l = needleIndices[l-1]
			}
		}
	}

	nInd := 0
	for i := 0; i < nh; {
		if haystack[i] == needle[nInd] {
			i++
			nInd++
			if nInd == nn {
				return i - nn
			}
		} else {
			if nInd == 0 {
				i++
			} else {
				nInd = needleIndices[nInd-1]
			}
		}
	}

	return -1
}

func main() {
	haystack, needle := "hello", "ll"

	index := strStr(haystack, needle)

	fmt.Println(index)
}

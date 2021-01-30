package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	ng := len(g)
	ns := len(s)

	if ng == 0 || ns == 0 {
		return 0
	}

	j := 0
	for i := 0; i < ns; i++ {
		if s[i] >= g[j] {
			j++
			if j == ng {
				break
			}
		}
	}
	return j
}

package leetcode

func hammingDistance(x int, y int) int {
	a := x ^ y

	count := 0
	for a > 0 {
		count++
		a &= a - 1
	}
	return count
}

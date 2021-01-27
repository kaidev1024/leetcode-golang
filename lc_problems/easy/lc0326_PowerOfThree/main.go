package leetcode

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for {
		if n%3 == 0 {
			n /= 3
		} else {
			break
		}
	}
	return n == 1
}

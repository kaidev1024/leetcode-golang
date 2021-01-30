package leetcode

func power(a int) int {
	result := 1
	for a > 0 {
		result <<= 1
		a--
	}
	return result
}

func findComplement(num int) int {
	count := 0
	copy := num
	for num > 0 {
		count++
		num >>= 1
	}

	n := power(count) - 1
	return n ^ copy
}

package leetcode

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	return addDigits(num%10 + addDigits(num/10))
}

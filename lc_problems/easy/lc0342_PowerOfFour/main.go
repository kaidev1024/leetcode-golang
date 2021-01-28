package leetcode

func isPowerOfFour(n int) bool {
	return n&0x5555555555555555 != 0 && n&(n-1) == 0
}

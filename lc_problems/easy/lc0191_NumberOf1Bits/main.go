package leetcode

func hammingWeight(num uint32) int {
	var n uint32 = 0
	for num > 0 {
		n += num & 1
		num >>= 1
	}
	return int(n)
}

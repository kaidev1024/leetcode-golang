package leetcode

func reverseBits(num uint32) uint32 {
	var r uint32 = 0

	for n := 32; n > 0; n-- {
		r <<= 1
		r = r | (num & 1)
		num >>= 1
	}

	return r
}

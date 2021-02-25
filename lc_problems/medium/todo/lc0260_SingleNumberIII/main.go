func singleNumber(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	// if xor is 0x0110, means num1^num2 equal to 0x0110
	// assume num1's second bit is 1, num2's second bit must be 0
	magic := 1
	for {
		if (xor & magic) == magic {
			break
		}
		magic <<= 1
	}

	num1, num2 := 0, 0
	for _, v := range nums {
		if (v & magic) == magic {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	return []int{num1, num2}
}
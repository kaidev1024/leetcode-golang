func singleNumber(nums []int) int {
	var ret int
	for i := 0; i < 64; i++ {
		sum, mask := 0, 1<<i
		for _, num := range nums {
			if num&mask != 0 {
				sum++
			}
		}
		if sum%3 != 0 {
			ret |= mask
		}
	}
	return ret
}
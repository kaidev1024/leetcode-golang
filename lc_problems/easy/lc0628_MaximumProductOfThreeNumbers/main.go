func maximumProduct(nums []int) int {
	l := len(nums)
	if l == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	sort.Ints(nums)
	p1 := nums[0] * nums[1] * nums[2]
	p2 := nums[0] * nums[1] * nums[l-1]
	p3 := nums[l-3] * nums[l-2] * nums[l-1]
	result := p1
	if p2 > result {
		result = p2
	}
	if p3 > result {
		return p3
	}
	return result
}
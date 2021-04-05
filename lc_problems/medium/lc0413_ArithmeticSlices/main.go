func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	dif := nums[1] - nums[0]
	count := 1
	ret := 0
	for i := 2; i < n; i++ {
		curDif := nums[i] - nums[i-1]
		if curDif == dif {
			ret += count
			count++
		} else {
			dif = curDif
			count = 1
		}
	}
	return ret
}
func findErrorNums(nums []int) []int {
	sum := 0
	dup := 0
	n := len(nums)

	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	for i := 0; i < n; i++ {
		v := abs(nums[i])
		sum += v
		ind := v - 1
		if nums[ind] < 0 {
			dup = v
		} else {
			nums[ind] = -nums[ind]
		}
	}

	result := make([]int, 2)
	result[0] = dup
	result[1] = (n+1)*n/2 - (sum - dup)

	return result
}
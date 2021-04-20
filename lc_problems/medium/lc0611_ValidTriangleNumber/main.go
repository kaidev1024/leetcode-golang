func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)

	res := 0

	for i := 2; i < len(nums); i++ {
		l, r := 0, i-1

		for l < r {
			if nums[l]+nums[r] > nums[i] {
				res += r - l
				r--
			} else {
				l++
			}
		}
	}

	return res
}
func nextPermutation(nums []int) {
	n := len(nums)
	flag := false
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] < nums[j] {
				flag = true
				nums[i], nums[j] = nums[j], nums[i]
				secondHalf := nums[i+1:]
				sort.Ints(secondHalf)
				nums = append(nums[:i+1], secondHalf...)
				break
			}
		}
		if flag {
			break
		}
	}

	if !flag {
		sort.Ints(nums)
	}
}
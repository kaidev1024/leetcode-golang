func checkPossibility(nums []int) bool {
	size := len(nums)

	ret := true
	for i := 1; i < size; i++ {
		//max_num = f(max_num, nums[i])
		if nums[i-1] > nums[i] {
			if i < 2 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
				break
				// times++
			} else {
				nums[i] = nums[i-1]
				break
				// times++
			}
		}
	}
	for j := 0; j < size-1; j++ {
		if nums[j] > nums[j+1] {
			ret = false
			break
		}
	}
	return ret
}
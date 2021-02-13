func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	first := 0
	second := -1

	for i := 1; i < n; i++ {
		if nums[i] > nums[first] {
			second = first
			first = i
			continue
		}
		if second == -1 || nums[i] > nums[second] {
			second = i
		}
	}
	if nums[first] >= 2*nums[second] {
		return first
	}
	return -1
}
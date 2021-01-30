package leetcode

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		pos := abs(nums[i]) - 1
		if nums[pos] > 0 {
			nums[pos] = -nums[pos]
		}
	}

	result := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

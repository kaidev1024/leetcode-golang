func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] < 0 {
			result = append(result, abs(num))
		} else {
			nums[index] = -nums[index]
		}
	}
	return result
}
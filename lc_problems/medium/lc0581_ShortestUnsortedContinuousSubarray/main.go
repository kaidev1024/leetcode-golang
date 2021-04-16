func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	start := 0
	for start < n-1 && nums[start+1] >= nums[start] {
		start++
	}
	if start == n-1 {
		return 0
	}

	end := n - 1
	for end > start && nums[end] >= nums[start] && nums[end] >= nums[end-1] {
		end--
	}
	min1 := nums[start]
	max1 := nums[start]
	for i := start; i <= end; i++ {
		min1 = min(min1, nums[i])
		max1 = max(max1, nums[i])
	}
	newStart := start
	for newStart >= 0 && min1 < nums[newStart] {
		newStart--
	}
	newEnd := end
	for newEnd < n && max1 > nums[newEnd] {
		newEnd++
	}
	return newEnd - newStart - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
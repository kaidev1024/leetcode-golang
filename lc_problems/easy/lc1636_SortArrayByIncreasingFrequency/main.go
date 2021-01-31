func frequencySort(nums []int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if counts[nums[i]] == counts[nums[j]] {
			return nums[i] > nums[j]
		}
		return counts[nums[i]] < counts[nums[j]]
	})

	return nums
}
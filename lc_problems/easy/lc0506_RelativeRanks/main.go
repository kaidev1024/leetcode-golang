func findRelativeRanks(nums []int) []string {
	l := len(nums)
	arr := make([]int, l)
	copy(arr, nums)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	m := make(map[int]int)

	for i, v := range arr {
		m[v] = i + 1
	}

	result := make([]string, l)

	for i := 0; i < l; i++ {
		v, _ := m[nums[i]]
		if v == 1 {
			result[i] = "Gold Medal"
		} else if v == 2 {
			result[i] = "Silver Medal"
		} else if v == 3 {
			result[i] = "Bronze Medal"
		} else {
			result[i] = strconv.Itoa(v)
		}
	}

	return result
}
func largestNumber(nums []int) string {

	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) > strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
	})

	var sb strings.Builder
	for i, n := 0, len(nums); i < n; i++ {
		sb.WriteString(strconv.Itoa(nums[i]))
	}

	pos := 0
	result := sb.String()
	n := len(result)

	for pos < n {
		if result[pos] == '0' {
			pos++
		} else {
			break
		}
	}
	if pos == n {
		return "0"
	}
	return result[pos:]
}
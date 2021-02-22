func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	nRemove := n - k

	result := make([]int, 0)
	for _, num := range nums {
		nr := len(result)
		for nr > 0 && num < result[nr-1] && nRemove > 0 {
			nr--
			result = result[:nr]
			nRemove--
		}
		result = append(result, num)
	}
	return result[:k]
}
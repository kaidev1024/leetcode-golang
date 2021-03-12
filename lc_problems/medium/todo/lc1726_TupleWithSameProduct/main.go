func tupleSameProduct(nums []int) int {
	counts := make(map[int]int)
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			counts[nums[i]*nums[j]]++
		}
	}
	result := 0
	for _, v := range counts {
		result += (v - 1) * v / 2
	}
	return result * 8
}
func subarraySum(nums []int, k int) int {
	dict := make(map[int]int)
	dict[0] = 1
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += dict[sum-k]
		dict[sum]++
	}
	return res
}
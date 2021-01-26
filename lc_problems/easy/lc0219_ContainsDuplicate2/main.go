package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	l := len(nums)
	set := make(map[int]struct{})
	for i := 0; i < k && i < l; i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = struct{}{}
	}
	for i := k; i < l; i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = struct{}{}
		delete(set, nums[i-k])
	}
	return false
}

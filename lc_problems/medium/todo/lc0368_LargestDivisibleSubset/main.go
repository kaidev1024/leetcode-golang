import "sort"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}
	sort.Ints(nums)
	parents := make([]int, n)
	counts := make([]int, n)
	maxCount := 0
	maxIndex := 0
	for start := n - 1; start >= 0; start-- {
		for end := start; end < n; end++ {
			if nums[end]%nums[start] == 0 && counts[end]+1 > counts[start] {
				counts[start] = counts[end] + 1
				parents[start] = end
				if counts[start] > maxCount {
					maxCount = counts[start]
					maxIndex = start
				}
			}
		}
	}
	ret := make([]int, 0)
	for {
		ret = append(ret, nums[maxIndex])
		if maxIndex == parents[maxIndex] {
			break
		}
		maxIndex = parents[maxIndex]
	}
	return ret
}


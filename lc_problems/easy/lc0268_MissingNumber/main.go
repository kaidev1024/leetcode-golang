package leetcode

func missingNumber(nums []int) int {
	l := len(nums)
	var sum int
	for _, val := range nums {
		sum += val
	}
	return l*(l+1)/2 - sum
}
